Rails.application.routes.draw do
  # For details on the DSL available within this file,
  # see https://guides.rubyonrails.org/routing.html
  resources :sparkles, only: [:index, :create]
  resources :users, only: [:show]
  root "sparkles#index"
end
