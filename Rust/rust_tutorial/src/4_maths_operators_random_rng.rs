#![allow(unused)]

use std::io;
use rand::Rng;
use std::io::{Write, BufReader, BufRead, ErrorKind};
use std::fs::File;
use std::cmp::Ordering;

fn main() {
    let num_1: f32 = 1.111111111111111;
    let num_2: f32 = 0.111111111111111;
    let num_3 = num_1 + num_2;
    println!("f32: {}", num_3);

    let num_4: f64 = 1.111111111111111;
    let num_5: f64 = 0.111111111111111;
    let num_6 = num_4 + num_5;
    println!("f64: {}", num_6);

    let mut num_7: u32 = 5;
    let num_8: u32 = 4;
    println!("5 % 4 = {}", num_7 % num_8);

    println!();

    let random_num = rand::thread_rng().gen_range(1..101);
    println!("Random: {}", random_num);
}