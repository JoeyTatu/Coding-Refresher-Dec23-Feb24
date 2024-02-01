#![allow(unused)]

use std::io;
use rand::Rng;
use std::io::{Write, BufReader, BufRead, ErrorKind};
use std::fs::File;
use std::cmp::Ordering;

fn main() {
    // Unsigned integer: u8, u16, u32, u64, u128, usize (positive values only)
    // Signed integer: i8, i16, i32, i64, i128, isize (minus and positive values)

    // Unsigned Integers
    println!("Unsigned integers:");
    println!("Min/Max u8: {}/{}", u8::MIN, u8::MAX);
    println!("Min/Max u16: {}/{}", u16::MIN, u16::MAX);
    println!("Min/Max u32: {}/{}", u32::MIN, u32::MAX);
    println!("Min/Max u64: {}/{}", u64::MIN, u64::MAX);
    println!("Min/Max u128: {}/{}", u128::MIN, u128::MAX);
    println!("Min/Max usize: {}/{}", usize::MIN, usize::MAX);

    // Signed Integers
    println!("\nSigned integers:");
    println!("Min/Max i8: {}/{}", i8::MIN, i8::MAX);
    println!("Min/Max i16: {}/{}", i16::MIN, i16::MAX);
    println!("Min/Max i32: {}/{}", i32::MIN, i32::MAX);
    println!("Min/Max i64: {}/{}", i64::MIN, i64::MAX);
    println!("Min/Max i128: {}/{}", i128::MIN, i128::MAX);
    println!("Min/Max isize: {}/{}", isize::MIN, isize::MAX);

    println!("\nFloating-point numbers:");
    println!("Min/Max f32: {}/{}", f32::MIN, f32::MAX);
    println!("Min/Max f64: {}/{}", f64::MIN, f64::MAX);
    
}