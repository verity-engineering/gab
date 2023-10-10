
#[derive(Debug)]
pub struct Config {
  server: String,
  user: String,
  password: String,
  repo: String,
  archive: String,
  releases: i32,
  prereleases: i32,
  retention: i32,
  debug: bool,
}

impl Config {
    pub fn build(mut args: impl Iterator<Item = String>) -> Result<Config, &'static str> {
        args.next();

        let server = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a server string"),
        };

        let user = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a user string"),
        };

        let password = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a password string"),
        };

        let repo = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a repo string"),
        };

        let archive = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a archive string"),
        };

        let releases = match args.next() {
            Some(arg) => { 
                let x = arg.parse::<i32>().unwrap();
                x
            },
            None => return Err("Didn't get a release string"),
        };

        let prereleases = match args.next() {
            Some(arg) => { 
                let x = arg.parse::<i32>().unwrap();
                x
            },
            None => return Err("Didn't get a prereleases string"),
        };

        let retention = match args.next() {
            Some(arg) => { 
                let x = arg.parse::<i32>().unwrap();
                x
            },
            None => return Err("Didn't get a retention string"),
        };

        let debug = match args.next() {
            Some(_) => true,
            None => return Err("Didn't get a debug flag"),
        };

        Ok(Config {
            server,
            user,
            password,
            repo,
            archive,
            releases,
            prereleases,
            retention,
            debug,
        })
    }
}
