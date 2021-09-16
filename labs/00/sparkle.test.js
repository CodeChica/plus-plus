const Sparkle = require("./sparkle");

test.skip("creates a new sparkle", () => {
  const sparkle = new Sparkle("@monalisa for helping me with my homework")

  expect(sparkle.sparklee).toBe("@monalisa")
  expect(sparkle.reason).toBe("for helping me with my homework")
});

test.skip("creates another sparkle", () => {
  const sparkle = new Sparkle("@kirstybrews for volunteering to help!")

  expect(sparkle.sparklee).toBe("@kirstybrews")
  expect(sparkle.reason).toBe("for volunteering to help!")
});
