require "minitest/autorun"

class Sparkle
  def self.create(sparklee, reason)
    # TODO:: Write your code here
  end
end

describe Sparkle do
  describe "when a visitor creates a sparkle" do
    before do
      @sparkles = []
      @sparkles.push(
        Sparkle.create("@monalisa", "for helping me with my homework!")
      )
    end

    it "gets added to the list of sparkles" do
      skip "it"

      assert_equal 1, @sparkles.count
      sparkle = @sparkles[0]
      assert_equal "@monalisa", sparkle.sparklee
      assert_equal "for helping me with my homework!", sparkle.reason
    end
  end
end
