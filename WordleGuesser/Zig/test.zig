const std = @import("std");

pub fn main() !void {
    const stdin = std.io.getStdIn().reader();
    const stdout = std.io.getStdOut().writer();
    var wordBuffer: [256]u8 = undefined;
    var checkBuffer: [256]u8 = undefined;
    //var currentGuess = [_]u8{ undefined, undefined, undefined, undefined, undefined };
    var checkGuess = [_]u8{ undefined, undefined, undefined, undefined, undefined };

    try stdout.print("Enter word: ", .{});

    const word = try stdin.readUntilDelimiter(&wordBuffer, '\n');

    try stdout.print("Enter validation: ", .{});

    const check = try stdin.readUntilDelimiter(&checkBuffer, '\n');

    for (check, 0..) |num, index| {
        checkGuess[index] = num;
    }

    std.debug.print("Word: {s}\n", .{word});
    std.debug.print("{u}\n", .{checkGuess});

    //for (checkGuess) |num| {
    //    switch (num) {
    //        0 => std.debug.print("Incorrect\n", .{}),
    //        1 => std.debug.print("Right letter, wrong place\n", .{}),
    //        2 => std.debug.print("Correct\n", .{}),
    //        else => std.debug.print("Not a valid entry\n", .{}),
    //    }
    //}

    //for (word) |letter| {
    //    //std.debug.print("{d}: {c}\t", .{ index, letter });
    //    std.debug.print("{c}\n", .{letter});
    //}
}
