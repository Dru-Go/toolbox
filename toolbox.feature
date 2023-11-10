Feature: Maintain a command line interface
  As a systems administrator
  I want to be able to migrate data from CSV to a database
  I want to recalculate the Transaction balance, unitPrice and cumulative (Weighted Mean)

  Scenario: Import Data from a csv file
    Given a CSV file
    Then I should read all Transaction from the file

  Scenario: Add Material Data to a Database
    Given a material name, category and Measurement
    Then I should add the material to the database

  Scenario: ReCalculate Transaction
    Given a MaterialId 
    Then I should scan all the Transaction records for the material and recalculate their prices from the last(old) Transaction to the most recent(new)
    Given a Transaction Id 
    Then I should scan all the Transactions after the TransactionID and recalculate their prices from the current Transaction to the most recent(new)