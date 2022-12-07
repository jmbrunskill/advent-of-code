use std::str::FromStr;
use std::cmp::Ordering;

/*
    This is an overkill use of structs and traits
    However, I thought it would be fun to implement the circular ordering of Rock, Paper, Scissors with the Ord trait
    Which caused me to get carried away...
*/


#[derive(Eq, PartialEq,PartialOrd, Debug)]
enum PaperRockScissors{
    Paper,
    Rock,
    Scissors,
}

impl FromStr for PaperRockScissors {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "A" => Ok(PaperRockScissors::Rock),
            "B" => Ok(PaperRockScissors::Paper),
            "C" => Ok(PaperRockScissors::Scissors),
            "X" => Ok(PaperRockScissors::Rock),
            "Y" => Ok(PaperRockScissors::Paper),
            "Z" => Ok(PaperRockScissors::Scissors),
            _ => Err(format!("Invalid game action {}", s)),
        }
    }
}

impl PaperRockScissors {
    fn shape_score(&self) -> i32 {
        match self {
            PaperRockScissors::Rock => 1,
            PaperRockScissors::Paper => 2,
            PaperRockScissors::Scissors => 3,
        }
    }
}

impl Ord for PaperRockScissors{
    fn cmp(&self, other: &Self) -> Ordering {
        match (self, other) {
            (PaperRockScissors::Paper, PaperRockScissors::Rock) =>Ordering::Greater,
            (PaperRockScissors::Rock, PaperRockScissors::Scissors) =>Ordering::Greater,
            (PaperRockScissors::Scissors, PaperRockScissors::Paper) =>Ordering::Greater,
            (PaperRockScissors::Paper, PaperRockScissors::Scissors) =>Ordering::Less,
            (PaperRockScissors::Rock, PaperRockScissors::Paper) =>Ordering::Less,
            (PaperRockScissors::Scissors, PaperRockScissors::Rock) =>Ordering::Less,
            (PaperRockScissors::Paper, PaperRockScissors::Paper) =>Ordering::Equal,
            (PaperRockScissors::Rock, PaperRockScissors::Rock) =>Ordering::Equal,
            (PaperRockScissors::Scissors, PaperRockScissors::Scissors) =>Ordering::Equal,
        }
    }
}

impl PaperRockScissors {
    fn win(&self) -> PaperRockScissors {
        match self {
            PaperRockScissors::Paper => PaperRockScissors::Scissors,
            PaperRockScissors::Rock => PaperRockScissors::Paper,
            PaperRockScissors::Scissors => PaperRockScissors::Rock,
        }
    }
    fn lose(&self) -> PaperRockScissors {
        match self {
            PaperRockScissors::Paper => PaperRockScissors::Rock,
            PaperRockScissors::Rock => PaperRockScissors::Scissors,
            PaperRockScissors::Scissors => PaperRockScissors::Paper,
        }
    }
    fn draw(&self) -> PaperRockScissors {
        match self {
            PaperRockScissors::Paper => PaperRockScissors::Paper,
            PaperRockScissors::Rock => PaperRockScissors::Rock,
            PaperRockScissors::Scissors => PaperRockScissors::Scissors,
        }
    }
}

#[derive(Debug)]
enum DesiredOutcome {
    Win,
    Lose,
    Draw,
}

#[derive(Debug)]
struct PaperRockScissorsRound {
    them: PaperRockScissors,
    us: PaperRockScissors,
    goal: DesiredOutcome,
}

impl FromStr for PaperRockScissorsRound {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut split = s.split_whitespace();
        let them = PaperRockScissors::from_str(split.next().unwrap()).unwrap();

        let us_str = split.next().unwrap();

        let us =  PaperRockScissors::from_str(us_str).unwrap();
        let goal = match us_str {
            "X" => DesiredOutcome::Lose,
            "Y" => DesiredOutcome::Draw,
            "Z" => DesiredOutcome::Win,
            _ => DesiredOutcome::Draw,
        };
        Ok(PaperRockScissorsRound { them, us,goal })
    }
}


impl PaperRockScissorsRound {
    fn score(&self) -> i32 {
        let shape_score = self.us.shape_score();
        let outcome_score = match self.us.cmp(&self.them) {
            Ordering::Greater => 6,
            Ordering::Equal => 3,
            Ordering::Less => 0,
        };

        shape_score + outcome_score
    }
}

/// In part one we need to calculate the expected score with X,Y,Z being our moves...
fn part1(input: &str) -> String {
    let mut total_score = 0;

    for line in input.lines(){
        let round = line.parse::<PaperRockScissorsRound>().unwrap();
        total_score += round.score();
        // println!("Total: {:?} - {:?} Round Score: {}", total_score, round, round.score());
    }

    return format!("{}", total_score);
}


// We need to calculate the expected score with X,Y,Z being the expected outcome
fn part2(input: &str) -> String {

    let mut total_score = 0;

    for line in input.lines(){
        let mut round = line.parse::<PaperRockScissorsRound>().unwrap();
        // println!("{:?}", round);
        round.us = match round.goal {
            DesiredOutcome::Win => round.them.win(),
            DesiredOutcome::Lose => round.them.lose(),
            DesiredOutcome::Draw => round.them.draw(),
        };
        total_score += round.score();
        // println!("Total: {:?} - {:?} Round Score: {}", total_score, round, round.score());
    }

    return format!("{}", total_score);
}

fn main() {
    // read input.txt into a string (this it totally not a prompt to get github copilot to help me...)
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
        let example1 = r#"A Y
B X
C Z"#;
        let expected1 = "15".to_string();

        assert_eq!(part1(example1), expected1);

        let example2 = r#"A Y
B Y
C Y"#;
        let expected2 = "15".to_string();

        assert_eq!(part1(example2), expected2);
    }

    #[test]
    fn test_part2() {
        let example1 = r#"A Y
B X
C Z"#;
        let expected1 = "12".to_string();

        assert_eq!(part2(example1), expected1);
    }
}