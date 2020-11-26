require 'test_helper'

class IntcodeControllerTest < ActionDispatch::IntegrationTest
  test "should get run" do
    get intcode_run_url
    assert_response :success
  end

end
