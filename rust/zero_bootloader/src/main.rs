// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

#![no_main]
#![no_std]

use uefi::{Status, entry, helpers};

pub(crate) mod panic;

#[entry]
fn main() -> Status {
    helpers::init();

    Status::SUCCESS
}
