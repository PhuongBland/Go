
const dropdownContainer = document.getElementById('dropdown-container')
const tableBody=document.getElementById('table-body')

const url = 'http://localhost:9000/api/categories'

const questionUrl="http://localhost:9000/api/question"
let questions=[];

function getLevel(number){
    if (number==1)
        return 'high';
    else if(number==2){
        return 'medium';

    }
    else return 'easy';

}
function getQuestionType(int){
    if (int==1)
        return 'text';
    else if(int ==2)
        return 'image';
    else return 'voice';
}

fetch(questionUrl)
    .then(response=>response.json())
    .then(result=>{
    if (result.status=="success"){
        questions=result.data;
    }
    })

function renderTable() {
    fetch(url)
        .then(response => response.json())
        .then(result => {

            const categoriesListModal = document.getElementById('list-categories');
            categoriesListModal.innerHTML = '';

            const status = result.status;

            // console.log('Data');
            // console.log(data);
            //
            // console.log('status');
            // console.log(status);

            if (status !== "success") {
                alert('Có lỗi trong quá trình lấy dữ liệu từ máy chủ');
                return;
            }

            const data = result.data;

            const parents = [];
            const children = [];

            for (const item of data) {
                if (item.parent_id === "") {
                    parents.push(item);
                } else {
                    children.push(item);
                }
            }

            for (const parent of parents) {
                const dropDownElement = document.createElement('div');
                dropDownElement.classList.add('dropdown');

                const dropButtonElement = document.createElement('button');
                dropButtonElement.classList.add('dropbtn');
                dropButtonElement.innerHTML = parent.name;

                //
                const dropdownContentElement = document.createElement('div');
                dropdownContentElement.classList.add('dropdown-content');

                dropDownElement.appendChild(dropButtonElement);
                dropDownElement.appendChild(dropdownContentElement);

                for (const child of children) {

                    categoriesListModal.insertAdjacentHTML('beforeend', `
                    <option value="${child['_id']}">${child.name}</option>
                `);

                    if (child.parent_id == parent._id) {
                        const aelement = document.createElement('a');
                        aelement.innerHTML = child.name;
                        aelement.style.cursor = 'pointer';

                        aelement.addEventListener('click', function () {
                            const category_ID = child._id;
                            tableBody.innerHTML = "";

                            for (const question of questions) {
                                if (question.category_id && question.category_id == category_ID) {
                                    const tr = document.createElement('tr');
                                    tr.innerHTML = `
                                                <td>${question.question}</td>
                                                <td>${child.name}</td>
                                                <td>${getQuestionType(question.content.type)}</td>
                                                <td>${getLevel(question.level)}</td>
                                                <td>${question.status}</td>
                                `;

                                    tableBody.appendChild(tr);
                                }
                            }
                        });

                        dropdownContentElement.appendChild(aelement);
                    }
                }

                dropdownContainer.appendChild(dropDownElement);
            }
        });
}

renderTable();

const btnAddQuestion = document.getElementById('add-question');
btnAddQuestion.addEventListener('click', () => {
    document.getElementById('new-question-modal').hidden = false;
});

document.getElementById('btn-close-modal').addEventListener('click', (e) => {
    e.preventDefault();
    document.getElementById('new-question-modal').hidden = true;
})

document.getElementById('btn-create-new-question').addEventListener('click', e => {
    e.preventDefault();

    console.log('taomoi');

    const cauHoi = document.getElementById('exampleInputEmail1').value;
    const allDapan = Array.from(document.getElementById('answer-list').querySelectorAll('input'));

    const dapan = [];

    allDapan.forEach(e => {
        dapan.push(e.innerHTML);
    });

    const dapAnDung = document.getElementById('dapandung').value;

    const nhomCauHoi = document.getElementById('list-categories').value;

    const capDo = document.getElementById('capdo').value;

    const kieuCauHoi = document.getElementById('cauhoioption').value;

    const status = 'active';

    const newCauHoi = {};

    newCauHoi.question = cauHoi;
    newCauHoi.level = capDo;
    newCauHoi.category_id = nhomCauHoi;
    newCauHoi.status = status;
    newCauHoi.content = {
        type: kieuCauHoi,
        content: cauHoi
    }
    newCauHoi.correct_codes = [dapAnDung];

    const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".split('');
    newCauHoi.answers = allDapan.map((a,idx) => {
        return {
            Code: characters[idx],
            Answer: a
        }
    })

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Cookie", "REVEL_FLASH=");

    var raw = JSON.stringify(newCauHoi);

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    fetch("http://localhost:9000/api/question/create", requestOptions)
        .then(response => response.text())
        .then(result => {
            console.log(result);
            document.getElementById('new-question-modal').hidden = true;
            renderTable();
        })
        .catch(error => console.log('error', error));
});