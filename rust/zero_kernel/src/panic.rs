// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

use core::panic::PanicInfo;
use uefi::println;

/// A custom panic handler, used across bootloader.
#[panic_handler]
pub fn panic_impl(info: &PanicInfo) -> ! {
    println!("panic {}", info);
    loop {}
}
