# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20160604013250) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "addresses", force: :cascade do |t|
    t.string   "phone_number"
    t.string   "address_line_1"
    t.string   "address_line_2"
    t.string   "city"
    t.string   "state"
    t.string   "country"
    t.string   "zip_code"
    t.datetime "created_at",     null: false
    t.datetime "updated_at",     null: false
  end

  create_table "job_profile_accomplishments", force: :cascade do |t|
    t.string   "name"
    t.integer  "job_profile_id"
    t.integer  "index"
    t.datetime "created_at",     null: false
    t.datetime "updated_at",     null: false
  end

  add_index "job_profile_accomplishments", ["job_profile_id"], name: "index_job_profile_accomplishments_on_job_profile_id", using: :btree

  create_table "job_profile_duties", force: :cascade do |t|
    t.string   "name"
    t.integer  "job_profile_id"
    t.integer  "index"
    t.datetime "created_at",     null: false
    t.datetime "updated_at",     null: false
  end

  add_index "job_profile_duties", ["job_profile_id"], name: "index_job_profile_duties_on_job_profile_id", using: :btree

  create_table "job_profile_skills", force: :cascade do |t|
    t.string   "name"
    t.integer  "job_profile_id"
    t.integer  "index"
    t.datetime "created_at",     null: false
    t.datetime "updated_at",     null: false
  end

  add_index "job_profile_skills", ["job_profile_id"], name: "index_job_profile_skills_on_job_profile_id", using: :btree

  create_table "job_profiles", force: :cascade do |t|
    t.string   "title"
    t.string   "employer_name"
    t.integer  "address_id"
    t.datetime "start_date"
    t.datetime "end_date"
    t.decimal  "salary"
    t.integer  "salary_type"
    t.integer  "average_weekly_hours"
    t.string   "supervisor_name"
    t.string   "supervisor_phone"
    t.datetime "created_at",           null: false
    t.datetime "updated_at",           null: false
    t.integer  "user_id"
  end

  add_index "job_profiles", ["address_id"], name: "index_job_profiles_on_address_id", using: :btree

  create_table "users", force: :cascade do |t|
    t.string   "email",                  default: "", null: false
    t.string   "encrypted_password",     default: "", null: false
    t.string   "reset_password_token"
    t.datetime "reset_password_sent_at"
    t.datetime "remember_created_at"
    t.integer  "sign_in_count",          default: 0,  null: false
    t.datetime "current_sign_in_at"
    t.datetime "last_sign_in_at"
    t.string   "current_sign_in_ip"
    t.string   "last_sign_in_ip"
    t.datetime "created_at",                          null: false
    t.datetime "updated_at",                          null: false
    t.string   "first_name",             default: "", null: false
    t.string   "last_name",              default: "", null: false
    t.string   "prefix"
    t.string   "suffix"
    t.string   "middle_initial"
    t.string   "website_url"
    t.integer  "address_id"
  end

  add_index "users", ["email"], name: "index_users_on_email", unique: true, using: :btree
  add_index "users", ["reset_password_token"], name: "index_users_on_reset_password_token", unique: true, using: :btree

  add_foreign_key "job_profile_accomplishments", "job_profiles"
  add_foreign_key "job_profile_duties", "job_profiles"
  add_foreign_key "job_profile_skills", "job_profiles"
  add_foreign_key "job_profiles", "addresses"
  add_foreign_key "job_profiles", "users"
end
