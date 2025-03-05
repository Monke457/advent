const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    const file = try std.fs.cwd().openFile("data/2015/day2.txt", .{});
    defer file.close();

    const reader = file.reader();
    var buf: [1024]u8 = undefined;

    var sum: i32 = 0;
    while (reader.readUntilDelimiterOrEof(&buf, '\n')) |maybe_line| {
        if (maybe_line == null) break;

        var it = std.mem.splitSequence(u8, maybe_line.?, "x");
        const l: i32 = try std.fmt.parseInt(i32, it.next().?, 10);
        const w: i32 = try std.fmt.parseInt(i32, it.next().?, 10);
        const h: i32 = try std.fmt.parseInt(i32, it.next().?, 10);

        const a1: i32 = w * l;
        const a2: i32 = w * h;
        const a3: i32 = h * l;

        const min: i32 = getMin(a1, a2, a3);

        const prod: i32 = 2 * (a1 + a2 + a3) + min;
        sum = sum + prod;
    } else |err| {
        print("Could not read file {}\n", .{err});
    }
    print("Needed paper: {d}\n", .{sum});
}

fn getMin(a1: i32, a2: i32, a3: i32) i32 {
    if (a1 < a2) {
        if (a1 < a3) {
            return a1;
        }
        return a3;
    }
    if (a2 < a3) {
        return a2;
    }
    return a3;
}
