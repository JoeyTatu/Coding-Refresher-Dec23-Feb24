#![allow(unused)]

use std::io;
use rand::Rng;
use std::io::{Write, BufReader, BufRead, ErrorKind};
use std::fs::File;
use std::cmp::Ordering;

fn main() {

    // if else
    let age = 8;
    if (age >= 1) && (age <= 18){
        println!("Important birthdate");
    } 
    
    else if (age == 21) || (age == 50){
        println!("Important birthdate");
    }

    else if (age >= 65){
        println!("Important Birthday!");
    }

    else {
        println!("Not an important birthday!");
    }

    // if else with bool
    let mut my_age = 47;
    let can_vote = if my_age >= 18{
        true
    }

    else {
        false
    };

    println!("Can Vote: {}", can_vote);

    // x matches y
    let age2 = 8;
    match age2 {
        1..=18 => println!("Age2: Important birthdate"),
        21 | 50 => println!("Age2: IImportant birthdate"),
        65..=i32::MAX => println!("Age2: IImportant birthdate"),
        _ => print!("Age2: Not an important birthdate"),
    };

    let my_age = 18;
    let voting_age = 18;

    match my_age.cmp(&voting_age){
        Ordering::Less => println!("Cannot vote"),
        Ordering::Greater => println!("Can vote"),
        Ordering::Equal => println!("Can vote from this year"),
    };

}


