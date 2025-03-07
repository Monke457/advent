const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    const file = try std.fs.cwd().openFile("data/2015/day3.txt", .{});
    defer file.close();

    var gpa = std.heap.page_allocator;

    const content = try file.readToEndAlloc(gpa, 1024 * 1024);
    defer gpa.free(content);

    var visited = std.AutoHashMap([2]i32, void).init(gpa);
    defer visited.deinit();

    var pos = [2]i32{ 0, 0 };
    try visited.put(pos, {});

    var posRobo = [2]i32{ 0, 0 };
    var isRobo: bool = false;

    for (content) |byte| {
        if (isRobo) {
            posRobo = updatePos(posRobo, byte);
            if (visited.contains(posRobo)) {} else {
                try visited.put(posRobo, {});
            }
        } else {
            pos = updatePos(pos, byte);
            if (visited.contains(pos)) {} else {
                try visited.put(pos, {});
            }
        }
        isRobo = !isRobo;
    }
    print("Houses visited at least once: {d}\n", .{visited.count()});
}

fn updatePos(pos: [2]i32, dir: u8) [2]i32 {
    var next = [2]i32{ pos[0], pos[1] };
    switch (dir) {
        '>' => next[1] += 1,
        '<' => next[1] -= 1,
        'v' => next[0] += 1,
        '^' => next[0] -= 1,
        else => return pos,
    }
    return next;
}
