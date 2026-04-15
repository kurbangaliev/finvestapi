// --- Функции работы с сервером ---

async function loadGroupNews(){
    const newsGroupContainer = document.getElementById('newsGroupContainer');
    try{
        const res = await fetch('https://localhost:8081/api/v1/groupNews');
        const newsGroupList = await res.json();
        newsGroupContainer.innerHTML='';
        newsGroupList.forEach((item,index)=>{
            const div=document.createElement('div'); div.className='newsGroup-item';

            const contentDiv=document.createElement('div'); contentDiv.className='container-sm';
            const title=document.createElement('input'); title.className='form-control'; title.value=item.title;

            const saveBtn=document.createElement('button'); saveBtn.className='btn btn-outline-primary'; saveBtn.textContent='Сохранить';
            saveBtn.onclick=()=>saveGroupNews(item.ID,title.value);

            const deleteBtn=document.createElement('button'); deleteBtn.className='btn btn-outline-danger'; deleteBtn.textContent='Удалить';
            deleteBtn.onclick=()=>deleteGroupNews(item.ID);

            contentDiv.appendChild(title);
            contentDiv.appendChild(saveBtn);
            contentDiv.appendChild(deleteBtn);

            div.appendChild(contentDiv);
            newsGroupContainer.appendChild(div);
        });
    }catch(err){ newsGroupContainer.innerHTML='<div>Ошибка загрузки новостей</div>'; }
}

async function saveGroupNews(id,title){
    try{
        const res=await fetch('https://localhost:8081/api/v1/groupNews/'+id,{
            method:'PUT',
            headers:{'Content-Type':'application/json'},
            body:JSON.stringify({id,title})
        });
        if (!res.ok) {
            const message = `An error occurred: ${res.status}`;
            throw new Error(message);
        }
        await loadGroupNews();
    }catch(error)
    {
        console.error('Fetch error:', error);
        alert('Ошибка при сохранении');
        throw error; // Re-throw if necessary
    }
}

async function deleteGroupNews(id){
    if(!confirm('Удалить группу?')) return;
    try{
        const res=await fetch('https://localhost:8081/api/v1/groupNews/'+id,{
            method:'DELETE',
            headers:{'Content-Type':'application/json'},
            body:JSON.stringify({id})
        });
        if(!res.ok) throw new Error();
        await loadGroupNews();
    }catch{ alert('Ошибка при удалении'); }
}

async function createGroupNews(){
    const submitBtn = document.getElementById('submitBtn');

    // Отправка группы новостей на сервер
    const title = document.getElementById('title').value.trim();

    if(!title) {
        alert('Заполните title');
        return;
    }

    const payload = { title };

    submitBtn.disabled = true;
    try {
        const res = await fetch('https://localhost:8081/api/v1/groupNews', {
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify(payload)
        });
        if(!res.ok) throw new Error();
        alert('Группа успешно добавлена');
        // Очистка формы
        document.getElementById('title').value = '';
    } catch(err) {
        alert('Ошибка при отправке группы на сервер');
    }
    submitBtn.disabled = false;
}
