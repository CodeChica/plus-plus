class SparklesController < ApplicationController
  def index
    @sparkles = Sparkle.recent.limit(50)
  end

  def create
  end
end
