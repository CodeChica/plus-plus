require "test_helper"

class UserTest < ActiveSupport::TestCase
   test "creating a new user" do
     assert User.create(handle: 'eilish')
   end

   test "cannot create a duplicate user" do
     assert_raises do
       User.create!(handle: 'billie')
     end
   end
end
