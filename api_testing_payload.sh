#!/bin/bash

echo "Adding Student: Aarav Sharma"
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '{"name": "Aarav Sharma", "age": 23, "course": "Mechanical Engineering", "roll": "R101"}'

echo "Adding Student: Ishika Verma"
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '{"name": "Ishika Verma", "age": 20, "course": "Computer Science", "roll": "R102"}'

echo "Adding Student: Kabir Singh"
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '{"name": "Kabir Singh", "age": 20, "course": "Electrical Engineering", "roll": "R103"}'

STUDENT_ID=1
echo "Deleting Student with ID: $STUDENT_ID"
curl -X DELETE http://localhost:8080/students/$STUDENT_ID

echo "Finding Student with ID: $STUDENT_ID"
curl -X GET http://localhost:8080/students/$STUDENT_ID

echo "Getting All Students"
curl -X GET http://localhost:8080/students

echo "Finding Students with Name: Aarav Sharma"
curl -X GET http://localhost:8080/students/name/Aarav%20Sharma

echo "Finding Students with Name: Ishika Verma"
curl -X GET http://localhost:8080/students/name/Ishika%20Verma

echo "Finding Student with Roll: R101"
curl -X GET http://localhost:8080/students/roll/R101

echo "Finding Student with Roll: R102"
curl -X GET http://localhost:8080/students/roll/R102

echo "Sorting Students by Age"
curl -X GET http://localhost:8080/students/sort/age
