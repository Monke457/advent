const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    const file = try std.fs.cwd().openFile("data/2015/day2.txt", .{});
    defer file.close();

    const reader = file.reader();
    var buf: [1024]u8 = undefined;

    var sum_paper: i32 = 0;
    var sum_ribbon: i64 = 0;
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
        const max: i32 = getMax(l, w, h);

        const prod_p: i32 = 2 * (a1 + a2 + a3) + min;
        const prod_r: i64 = 2 * (l + h + w - max) + (l * h * w);

        sum_paper = sum_paper + prod_p;
        sum_ribbon = sum_ribbon + prod_r;
    } else |err| {
        print("Could not read file {}\n", .{err});
    }
    print("Paper needed: {d}\n", .{sum_paper});
    print("Ribbon needed: {d}\n", .{sum_ribbon});
}

fn getMax(a1: i32, a2: i32, a3: i32) i32 {
    if (a1 > a2) {
        if (a1 > a3) {
            return a1;
        }
        return a3;
    }
    if (a2 > a3) {
        return a2;
    }
    return a3;
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
