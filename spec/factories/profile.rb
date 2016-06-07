FactoryGirl.define do
  factory :job_profile do
    title {Faker::Name.title}
    employer_name {Faker::App.name}
    start_date {Faker::Date.between(10.years.ago, Date.current)}
    end_date {Faker::Date.between(start_date + 2.months, Date.current)}
    salary 100000
    salary_type 'yearly'
    average_weekly_hours 40
    supervisor_name {Faker::Name.name}
    supervisor_phone {Faker::PhoneNumber.phone_number}
    address

    factory :job_profile_with_notes do
      transient do
          duty_count 3
          skill_count 3
          accomplishment_count 3
      end

      after(:create) do |job, evaluator|
        create_list(:duty, evaluator.duty_count, job_profile: job)
      end
      after(:create) do |job, evaluator|
        create_list(:skill, evaluator.skill_count, job_profile: job)
      end
      after(:create) do |job, evaluator|
        create_list(:accomplishment, evaluator.accomplishment_count, job_profile: job)
      end
    end
  end


  factory :job_profile_duty, aliases: [:duty] do
    name {Faker::Commerce.product_name + ' ' + (%w{Stocking Cleaning Exporting Importing}.sample)}
  end

  factory :job_profile_skill, aliases: [:skill] do
    name {Faker::Commerce.product_name + ' ' + (%w{Demos Marketing Sales Rendering}.sample)}
  end

  factory :job_profile_accomplishment, aliases: [:accomplishment] do
    name {Faker::Commerce.department + ' ' + (["Sales up #{Faker::Number.decimal(2, 0)}%"].sample)}
  end
end