// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

#![no_main]
#![no_std]

pub(crate) mod core;
pub(crate) mod panic;

#[cfg(target_arch = "x86_64")]
#[unsafe(no_mangle)]
pub extern "win64" fn efi_main() -> ! {
    core::entry();
    loop {}
}

#[cfg(target_arch = "aarch64")]
#[unsafe(no_mangle)]
pub extern "C" fn efi_main() -> ! {
    core::entry();
    loop {}
}
