
// This tool is used to:
// 1.) connect to artifactory
// 2.) get a list of artifacts in a repository
// 3.) find duplicates and clear them up
// 4.) move binaries that are out of place. should be <repo>/<artifactname>/<artifactory><version>
// 5.) archive anything within a set verions
//       <M>.<m>.<b>  
// 6.) delete anything older than x months in the archive
use crate::config::Config;
use std::env;
use std::process;

mod config;

fn main() {
  let args: Vec<String> = env::args().collect();
  let _config = Config::build(args.into_iter()).unwrap_or_else(|err| {  
        eprintln!("Problem parsing arguments: {err}");
        process::exit(1);});
}
