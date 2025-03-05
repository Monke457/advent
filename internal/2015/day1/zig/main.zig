const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    const file = try std.fs.cwd().openFile("data/2015/day1.txt", .{});
    defer file.close();
    const reader = file.reader();

    var pos: i32 = 0;
    var basement: i32 = 0;
    var i: i32 = 0;

    while (reader.readByte()) |byte| {
        i = i + 1;
        if (byte == 10) break;
        if (byte == 40) pos = pos + 1;
        if (byte == 41) pos = pos - 1;
        if (pos < 0 and basement == 0) {
            basement = i;
        }
    } else |err| {
        print("Error while reading the file: {}\n", .{err});
    }

    print("final position: {d}\n", .{pos});
    print("enters basement: {d}\n", .{basement});
}
