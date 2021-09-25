require "minitest/autorun"

class Sparkle
  attr_reader :sparklee, :reason

  def initialize(sparklee, reason)
    # TODO::
  end
end

describe Sparkle do
  describe "when a sparkle is created" do
    before do
      @recipient = "@monalisa"
      @reason = "for helping me with my homework!"

      @sparkle = Sparkle.new(@recipient, @reason)
    end

    it "saves the recipient" do
      skip "it"

      assert_equal @recipient, @sparkle.sparklee
    end

    it "saves the reason" do
      skip "it"

      assert_equal @reason, @sparkle.reason
    end
  end
end
