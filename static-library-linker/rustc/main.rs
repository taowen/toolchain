#[link(name = "file1")]
extern "C" {
    fn hello() -> i32;
}

#[link(name = "file2")]
extern "C" {
    fn world() -> i32;
}

fn main() {
    let r = unsafe { hello() + world() };
    ::std::process::exit(r)
}