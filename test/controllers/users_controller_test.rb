require "test_helper"

class UsersControllerTest < ActionDispatch::IntegrationTest
  setup do
    @user = users(:billie)
  end

  test "should show user" do
    get user_url(@user)
    assert_response :success
  end
end
