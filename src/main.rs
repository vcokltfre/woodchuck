use iter_tools::Itertools;
use std::env;

fn main() {
    if let Some((_, full, start, end)) = env::args().collect_tuple() {
        let qualifier = match full.chars().nth(0).unwrap() {
            'a' | 'e' | 'i' | 'o' | 'u' => "an",
            _ => "a",
        };

        println!(
            "How much {} could {} {} {} if {} {} could {} {}?",
            start, qualifier, full, end, qualifier, full, end, start
        );
    } else {
        println!("Usage: woodchuck <full> <start> <end>");
    }
}