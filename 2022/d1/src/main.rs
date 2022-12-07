use log;
use std::collections::BinaryHeap;


/// In part one we need to figure out which elf is carrying the most food, elves are separated by a blank line
fn part1(input: &str) -> String {

    let mut max_food = 0;
    let mut current_elf_food = 0;

    // split the input by lines
    for line in input.lines(){
        log::debug!("{}", line);
        if line == ""{
            // if we have a blank line, we have reached the end of an elf's food
            current_elf_food = 0;
            continue;
        }
        current_elf_food += line.parse::<i32>().unwrap();
        if current_elf_food > max_food {
           max_food = current_elf_food;
        }
    }

    return format!("{}", max_food);
}


/// In part 2 we need to figure out the total food held by the top 3 elves.
fn part2(input: &str) -> String {


    let mut heap_of_food = BinaryHeap::new();

    let mut current_elf_food = 0;

    // split the input by lines
    for line in input.lines(){
        log::debug!("{}", line);
        if line == ""{
            // if we have a blank line, we have reached the end of an elf's food
            heap_of_food.push(current_elf_food);
            current_elf_food = 0;
            continue;
        }
        current_elf_food += line.parse::<i32>().unwrap();
    }

    let top3_food = heap_of_food.pop().unwrap() + heap_of_food.pop().unwrap() + heap_of_food.pop().unwrap();

    return format!("{}", top3_food);
}

fn main() {
    env_logger::init();
    // read input.txt into a string
    let input = std::fs::read_to_string("input.txt").expect("Unable to read input file");

    println!("Part1: {}", part1(&input));
    println!("Part2: {}", part2(&input));
}


#[cfg(test)]
mod tests {
    use crate::part1;
    use crate::part2;

    #[test]
    fn test_part1() {
        let example1 = r#"1000
2000
3000

4000

5000
6000

7000
8000
9000

10000"#;
        let expected1 = "24000".to_string();

        assert_eq!(part1(example1), expected1);
    }

    #[test]
    fn test_part2() {
        let example1 = r#"1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

"#;//Added blank line on the end to match input
        let expected1 = "45000".to_string();

        assert_eq!(part2(example1), expected1);
    }
}