# =============================
# Settings
# =============================
# Tools
ISOGEN := xorrisofs

# Directories
BUILD := build
ISO_DIR := $(BUILD)/iso

# Options
BUILD_TYPE ?= debug
ARCH ?= x86_64

# Decide bootloader file
BOOTLOADER_FILE := BOOTX64.EFI
ifeq ($(ARCH), aarch64)
	BOOTLOADER_FILE = BOOTA64.EFI
endif
ifeq ($(ARCH), x86_64)
	BOOTLOADER_FILE := BOOTX64.EFI
endif

# To get EFI files, compiled from Rust, we need to get if from rust/target directory,
# and rename them.
KERNEL_EFI_COMPILED := rust/target/$(ARCH)-unknown-uefi/$(BUILD_TYPE)/zero_kernel.efi
KERNEL_EFI_TOTAL := $(ISO_DIR)/EFI/BOOT/zero_kernel.efi
BOOTLOADER_EFI_COMPILED := rust/target/$(ARCH)-unknown-uefi/$(BUILD_TYPE)/zero_bootloader.efi
BOOTLOADER_EFI_TOTAL := $(ISO_DIR)/EFI/BOOT/$(BOOTLOADER_FILE)

# Cargo flags
CARGO_FLAGS := --target $(ARCH)-unknown-uefi
ifeq ($(BUILD_TYPE), release)
	CARGO_FLAGS += --release
endif

# EFI image for ISO
EFI_IMG := $(BUILD)/efi.img

# Final target
TARGET := $(BUILD)/zero.iso

# QEMU Emulation
OVMF ?=
QEMU := qemu-system-$(ARCH)
QEMU_FLAGS =
ifeq ($(ARCH), x86_64)
	QEMU_FLAGS := \
		-drive file=$(TARGET),media=cdrom \
		-m 2G \
		-drive if=pflash,format=raw,readonly=on,file=$(OVMF) \
		-boot d \
		-net none \
		-serial stdio
endif
ifeq ($(ARCH), aarch64)
	QEMU_FLAGS := \
		-machine virt,secure=off \
		-drive file=$(TARGET),media=cdrom \
		-cpu cortex-a57 \
		-drive if=pflash,format=raw,readonly=on,file=$(OVMF) \
		-m 2G \
		-nographic
endif

# =============================
# Jobs
# =============================
all: clean $(TARGET)

# Create EFI FAT image
$(EFI_IMG): $(KERNEL_EFI_TOTAL) $(BOOTLOADER_EFI_TOTAL)
	dd if=/dev/zero of=$(EFI_IMG) bs=1M count=10
	mkfs.vfat -F 12 $(EFI_IMG)
	mmd -i $(EFI_IMG) ::/EFI
	mmd -i $(EFI_IMG) ::/EFI/BOOT
	mcopy -i $(EFI_IMG) $(BOOTLOADER_EFI_TOTAL) ::/EFI/BOOT/
	mcopy -i $(EFI_IMG) $(KERNEL_EFI_TOTAL) ::/EFI/BOOT/

# Main target - create ISO with EFI image
$(TARGET): $(EFI_IMG)
	$(ISOGEN) -J -R -V "Zero" \
		-eltorito-alt-boot \
		-e efi.img \
		-no-emul-boot \
		-isohybrid-gpt-basdat \
		-o $(TARGET) \
		$(EFI_IMG)

# Directories
$(EFI_DIR): $(BUILD)
	mkdir $(EFI_DIR) -p
$(ISO_DIR): $(BUILD)
	mkdir $(ISO_DIR) -p
	mkdir $(ISO_DIR)/EFI
	mkdir $(ISO_DIR)/EFI/BOOT

# Compile kernel/bootloader, putting inside EFI_DIR folder
__cargo:
	cd rust && cargo build $(CARGO_FLAGS)
$(KERNEL_EFI_TOTAL): __cargo $(ISO_DIR)
	mv $(KERNEL_EFI_COMPILED) $(KERNEL_EFI_TOTAL)
$(BOOTLOADER_EFI_TOTAL): __cargo $(ISO_DIR)
	mv $(BOOTLOADER_EFI_COMPILED) $(BOOTLOADER_EFI_TOTAL)

$(BUILD):
	mkdir $(BUILD) -p

clean:
	rm $(BUILD) -fr

# Run ISO file in qemu
qemu: clean $(TARGET)
	$(QEMU) $(QEMU_FLAGS)