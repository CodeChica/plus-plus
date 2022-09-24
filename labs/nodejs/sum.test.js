const sum = require("./sum");

test("adds 1 + 2 to equal 3", () => {
  expect(sum(1, 2)).toBe(3);
});

test("adds 100 + 20 to equal 120", () => {
  expect(sum(100, 20)).toBe(120);
});

test("adds 12 + 7 to equal 19", () => {
  expect(sum(12, 7)).toBe(19);
});
