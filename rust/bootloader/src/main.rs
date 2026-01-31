// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

#![no_main]
#![no_std]

pub(crate) mod panic;

use core::time::Duration;
use uefi::{Status, entry, println};

#[entry]
fn main() -> Status {
    uefi::helpers::init().unwrap();
    println!("Hi :>");
    uefi::boot::stall(Duration::from_secs(10));
    Status::SUCCESS
}
