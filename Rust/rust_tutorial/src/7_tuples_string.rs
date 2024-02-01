#![allow(unused)]

use std::io;
use rand::Rng;
use std::io::{Write, BufReader, BufRead, ErrorKind};
use std::fs::File;
use std::cmp::Ordering;

fn main() {
    let my_tuple: (u8, String, f64) = (47, "Derek".to_string(), 50_000.00);

    println!("Name : {}", my_tuple.1);

    let(v1, v2, v3) = my_tuple;
    println!("Age : {}", v1);

    ////////////////////
    
    // String - can be changed 
    // &str - reference, point to string, allow viewing

    let mut st1 = String::new();
    st1.push('A');
    st1.push_str(" word");

    for word in st1.split_whitespace(){
        println!("{}", word)
    }

    let st2 = st1.replace("A", "Another");
    println!("{}", st2);

    //////
    
    let st3 = String::from("x r t b h k k a m c");
    let mut v1: Vec<char> = st3.chars().collect();
    v1.sort();
    v1.dedup();

    for char in v1 {
        println!("{}", char);
    }

    //////
    
    let st4: &str = "Random string";
    let mut st5: String = st4.to_string();
    println!("{}", st5);

    let byte_arr1 = st5.as_bytes();
    let st6 = &st5[0..6];
    println!("String len : {}", st6.len());

    st5.clear();

    let st6 = String::from("Just some");
    let st7 = String::from(" words");
    let st8 = st6 + &st7;
    
    for char in st8.bytes(){
        println!("{}", char)
    }


}


