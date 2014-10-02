require 'capybara'
require 'capybara/dsl'
    
Capybara.default_driver = :selenium
Capybara.app_host = 'http://localhost:8080'

RSpec.configure do |config|
  config.include Capybara::DSL
end

describe 'home page' do
  it 'allows you to see blog posts' do
    visit '/'
    page.should have_content('Finding The Results Of A Batch Upsert With MySQL')
  end

  it 'allows you to click on blog posts' do
    visit '/'
    click_on 'Finding The Results Of A Batch Upsert With MySQL'
    page.should have_content('So, how would we go about getting the IDs of those that were inserted and those that were updated?')
  end
end
