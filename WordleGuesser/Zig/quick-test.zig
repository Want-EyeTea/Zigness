const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    const stdout = std.io.getStdOut().writer();
    var buf: [256]u8 = undefined;
    const bgn = std.heap.page_allocator;
    var array = std.ArrayList(u8).init(bgn);

    try stdout.print(":>> ", .{});

    const input = try stdin.readUntilDelimiter(&buf, '\n');

    // FOR DEBUGGING
    //std.debug.print("----------", .{});
    //std.debug.print("----------\n", .{});

    for (input) |num| {
        if (std.ascii.isDigit(num)) {
            try array.append(@intCast(num - '0'));
        }
    }

    try validateGuess(array);

    array.deinit();
}

fn validateGuess(vInput: std.ArrayList(u8)) !void {
    for (vInput.items) |a| {
        switch (a) {
            0 => std.debug.print("Incorrect\n", .{}),
            1 => std.debug.print("Close\n", .{}),
            2 => std.debug.print("Correct\n", .{}),
            else => std.debug.print("Invalid Entry\n", .{}),
        }
    }
}
