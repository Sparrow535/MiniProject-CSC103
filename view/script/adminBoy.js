window.onload = function () {
    fetch('/students')
        .then(response => response.text())
        .then(data => showStudents(data))
        .catch(error => console.error('Error fetching students:', error));
};

function updateSecondSelect() {
    var firstSelect = document.getElementById("block");
    var secondSelect = document.getElementById("secondSelect");
    var floor = document.getElementById("floor");

    // Clear existing options
    secondSelect.innerHTML = "";

    // Add options based on the selected value of the first select
    if (firstSelect.value === "K") {
        var blockLetter = "K";
    } else if (firstSelect.value === "L") {
        var blockLetter = "L";
    } else if (firstSelect.value === "M") {
        var blockLetter = "M";
    } else if (firstSelect.value === "N") {
        var blockLetter = "N";
    }

    if (floor.value === "Fo") {
        var floorLetter = "Fo";
    } else if (floor.value === "T") {
        var floorLetter = "T";
    } else if (floor.value === "S") {
        var floorLetter = "S";
    } else if (floor.value === "F") {
        var floorLetter = "F";
    }

    if (blockLetter && floorLetter) {
        for (let i = 1; i <= 8; i++) {
            let option = document.createElement("option");
            option.value = String(blockLetter + floorLetter + "-" + i);
            option.text = blockLetter + floorLetter + "-" + i;
            secondSelect.add(option);
        }
    }
}

function filterStudentsByBlock() {
    var selectedBlock = document.getElementById("blocks").value;
    var table = document.getElementById("myTable");

    // Clear all rows except the header
    for (var i = table.rows.length - 1; i > 0; i--) {
        table.deleteRow(i);
    }

    fetch('/students')
        .then(response => response.json())
        .then(data => {
            console.log(data); // Log the data to ensure it contains student information

            // Filter students by selected block
            const filteredStudents = data.filter(student => {
                const isMatch = selectedBlock === "0" || student.room.startsWith(selectedBlock);
                console.log(`${student.room} - ${selectedBlock} - ${isMatch}`); // Log for debugging
                return isMatch;
            });

            showStudents(JSON.stringify(filteredStudents)); // Use showStudents to display filtered results
        })
        .catch(error => console.error('Error fetching students:', error));
}


function addStudent() {
    var data = getFormData();

    fetch('/student', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json; charset=UTF-8'
        },
        body: JSON.stringify(data)
    }).then(response => {
        if (response.ok) {
            fetch('/student/' + data.enrollment_no)
                .then(response => response.text())
                .then(data => showStudent(data));
        } else {
            throw new Error(response.status);
        }
    }).catch(e => {
        if (e.message == 303) {
            alert("User not logged in.");
            window.open("index.html", "_self");
        } else if (e.message == 500) {
            alert("Server error!");
        } else {
            alert("Error: " + e.message);
        }
    });
    resetForm();
}


function resetForm() {
    document.getElementById('fname').value = '';
    document.getElementById('lname').value = '';
    document.getElementById('enroll').value = '';
    document.getElementById('contact').value = '';
    document.getElementById('block').value = '0';
    document.getElementById('floor').value = '0';
    document.getElementById('secondSelect').innerHTML = '<option value="0">Select Room</option>';
}

function showStudent(data) {
    const student = JSON.parse(data)
    newRow(student);
}

function showStudents(data) {
    const students = JSON.parse(data);
    students.forEach(student => {
        newRow(student);
    });
}
var selectedRow = null;

function updateStudent(r) {
    selectedRow = r.parentElement.parentElement;
    document.getElementById('fname').value = selectedRow.cells[0].innerHTML;
    document.getElementById('lname').value = selectedRow.cells[1].innerHTML;
    document.getElementById('enroll').value = selectedRow.cells[2].innerHTML;
    document.getElementById('contact').value = selectedRow.cells[3].innerHTML;

    // Extract room number from the table cell
    var roomNo = selectedRow.cells[4].innerHTML;

    // Fetch complete room details including block and floor information
    fetch('/room/' + roomNo)
        .then(response => response.json())
        .then(data => {
            // Populate form fields with room details
            // Set the selected block and floor
            var blockDropdown = document.getElementById('block');
            var floorDropdown = document.getElementById('floor');

            blockDropdown.value = data.block;
            floorDropdown.value = data.floor;

            // Update the second select options
            updateSecondSelect();

            // Set the selected room number
            var roomNumberDropdown = document.getElementById('secondSelect');
            for (var k = 0; k < roomNumberDropdown.options.length; k++) {
                if (roomNumberDropdown.options[k].value.trim() === data.room_no.trim()) {
                    roomNumberDropdown.selectedIndex = k;
                    break;
                }
            }
        })
        .catch(error => console.error('Error fetching room details:', error))

    var btn = document.getElementById("addBTN");
    var enrollment_no = selectedRow.cells[2].innerHTML;
    if (enrollment_no) {
        btn.innerHTML = "Update";
        // Pass enrollment_no to update function
        btn.setAttribute("onclick", "update(" + enrollment_no + ")");
    }
}


function update(enrollment_no) {
    var newData = getFormData();
    fetch('/student/' + enrollment_no, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json; charset=UTF-8'
        },
        body: JSON.stringify(newData)
    })
        .then(response => {
            if (response.ok) {
                // Update table cell with new data
                selectedRow.cells[0].innerHTML = newData.first_name;
                selectedRow.cells[1].innerHTML = newData.last_name;
                selectedRow.cells[2].innerHTML = newData.enrollment_no;
                selectedRow.cells[3].innerHTML = newData.contact_no;
                selectedRow.cells[4].innerHTML = newData.room;

                var button = document.getElementById("addBTN");
                button.innerHTML = "Add";
                button.setAttribute("onclick", "addStudent()");
                selectedRow = null;
                resetForm();
            } else {
                alert("Server: Update request error");
            }
        })
        .catch(error => console.error('Error updating student:', error));
}


function deleteStudent(r) {
    if (confirm("Are you sure you want to delete this?")) {
        selectedRow = r.parentElement.parentElement;
        var enrollment_no = selectedRow.cells[2].innerHTML;

        fetch('/student/' + enrollment_no, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json; charset=UTF-8'
            }
        })
            .then(response => {
                if (response.ok) {
                    var rowIndex = selectedRow.rowIndex;
                    if (rowIndex > 0) {
                        document.getElementById("myTable").deleteRow(rowIndex);
                    }
                    selectedRow = null;
                } else {
                    alert("Server: Delete request error");
                }
            })
            .catch(error => console.error('Error deleting student:', error));
    }
}

function newRow(student) {
    var table = document.getElementById('myTable');
    var row = table.insertRow(table.length);

    var td = [];
    for (var i = 0; i < table.rows[0].cells.length; i++) {
        td[i] = row.insertCell(i);
    }
    td[0].innerHTML = student.first_name;
    td[1].innerHTML = student.last_name;
    td[2].innerHTML = student.enrollment_no;
    td[3].innerHTML = student.contact_no;
    td[4].innerHTML = student.room;
    td[5].innerHTML = '<input type="button" onclick="updateStudent(this)" value="Edit" id="editStudent">';
    td[6].innerHTML = '<input type="button" onclick="deleteStudent(this)" value="Delete" id="deleteStudent">';
}

function getFormData() {
    return {
        first_name: document.getElementById('fname').value,
        last_name: document.getElementById('lname').value,
        enrollment_no: parseInt(document.getElementById('enroll').value),
        contact_no: parseInt(document.getElementById('contact').value),
        room: document.getElementById('secondSelect').value
    };
}
