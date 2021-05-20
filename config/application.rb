require_relative "boot"

require "rails"
require "active_model/railtie"
require "active_record/railtie"
require "action_controller/railtie"
require "action_view/railtie"
require "rails/test_unit/railtie"
require "primer/view_components/engine"

Bundler.require(*Rails.groups)

module Sparkles
  class Application < Rails::Application
    config.load_defaults 6.1
    config.generators.system_tests = nil
    config.logger = Logger.new(STDOUT)
  end
end
