#![allow(unused)]

use std::io;
use rand::Rng;
use std::io::{Write, BufReader, BufRead, ErrorKind};
use std::fs::File;
use std::cmp::Ordering;

fn main() {
    let int_u8: u8 = 5;
    let int2_u8: u8 = 4;

    let int3_u32: u32 = (int_u8 as u32) + (int2_u8 as u32);

    ////
    
    enum Day{
        Monday,
        Tuesday,
        Wednesday,
        Thursday,
        Friday,
        Saturday,
        Sunday,
    }

    impl Day {
        fn is_weekend(&self) -> bool{
            match self{
                Day::Saturday | Day::Sunday => true,
                _ => false,
            }
        }
    }

    let today:Day = Day::Monday;

    match today {
        Day::Monday => println!("Everyone hates Mondays"),
        Day::Tuesday => println!("It's Tuesday!"),
        Day::Wednesday => println!("Hump day!"),
        Day::Thursday => println!("Almost there, it's Thursday!"),
        Day::Friday => println!("TGIF!"),
        Day::Saturday => println!("Enjoy the weekend!"),
        Day::Sunday => println!("Lazy Sunday"),
    }

    println!("Is today the weekend? {}", today.is_weekend());

}
    