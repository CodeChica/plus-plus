require "minitest/autorun"

class Sparkle
  attr_reader :recipient, :reason

  def initialize(recipient, reason)
    @recipient = recipient
    @reason = reason
  end
end

describe Sparkle do
  describe "when a @monalisa is Sparkled" do
    before do
      @recipient = "@monalisa"
      @reason = "for helping me with my homework!"

      @sparkle = Sparkle.new(@recipient, @reason)
    end

    it "saves the recipient" do
      assert_equal @recipient, @sparkle.recipient
    end

    it "saves the reason" do
      assert_equal @reason, @sparkle.reason
    end
  end

  describe "when @GabbyGap is Sparkled" do
    before do
      # arrange
      @recipient = "@GabbyGap"
      @reason = "for volunteering to share her screen!"

      # act
      @sparkle = Sparkle.new(@recipient, @reason)
    end

    it "saves username of the recipient" do
      assert_equal "@GabbyGap", @sparkle.recipient
    end

    it "saves the reason for the sparkle" do
      assert_equal @reason, @sparkle.reason
    end
  end
end
