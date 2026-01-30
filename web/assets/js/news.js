// --- Функции работы с сервером ---
function fillSelectType(selectElement){
    let option = document.createElement('option');
    option.value = "1";
    option.textContent = "YouTube";
    selectElement.appendChild(option);

    option = document.createElement('option');
    option.value = "2";
    option.textContent = "Image";
    selectElement.appendChild(option);

    option = document.createElement('option');
    option.value = "3";
    option.textContent = "Pdf";
    selectElement.appendChild(option);

    option = document.createElement('option');
    option.value = "4";
    option.textContent = "Text";
    selectElement.appendChild(option);
}

async function loadNews(){
    const newsContainer = document.getElementById('newsContainer');
    try{
        const dislike = 0;
        const like = 1;
        const userId = 0;
        const res = await fetch('https://localhost:8081/api/v1/news/');
        const newsList = await res.json();
        newsContainer.innerHTML='';
        newsList.forEach((item,index)=>{
            const div=document.createElement('div'); div.className='news-item';

            const contentDiv=document.createElement('div'); contentDiv.className='container-sm';
            const title=document.createElement('input'); title.className='form-control'; title.value=item.title;
            const messageDate=document.createElement('input'); messageDate.className='form-control'; messageDate.type='date'; messageDate.value=item.messageDate.split('T')[0];
            const message=document.createElement('textarea'); message.className='form-control'; message.value=item.message;

            const type=document.createElement('select'); type.className='form-control'; fillSelectType(type); type.value=item.type;

            const mediaLink=document.createElement('input'); mediaLink.className='form-control'; mediaLink.value=item.mediaLink;
            const downloadLink=document.createElement('input'); downloadLink.className='form-control'; downloadLink.value=item.downloadLink;

            const statusId=document.createElement('select'); statusId.className='form-control'; fillSelectStatus(statusId); statusId.value=item.statusId;

            const newsId=document.createElement('input'); newsId.className='form-control'; newsId.value=item.ID; newsId.disabled=true;

            const saveBtn=document.createElement('button'); saveBtn.className='btn btn-outline-primary'; saveBtn.textContent='Сохранить';
            saveBtn.onclick=()=>saveNews(item.ID,title.value,messageDate.value,message.value,type.value,mediaLink.value,
                downloadLink.value,statusId.value,item.authorId,item.liked,item.disliked,item.viewCount);

            const deleteBtn=document.createElement('button'); deleteBtn.className='btn btn-outline-danger'; deleteBtn.textContent='Удалить';
            deleteBtn.onclick=()=>deleteNews(item.ID);

            const divLiked=document.createElement('div'); div.className='news-item';
            let space = document.createTextNode(" ");
            const likeBtn=document.createElement('img'); likeBtn.src="assets/icons/like.svg"; likeBtn.className='news-like-sm zoom-image';
            likeBtn.onclick=()=>likeNews(item.ID, userId, like);
            const spanLiked=document.createElement('span'); spanLiked.className='badge text-bg-secondary';spanLiked.textContent=item.liked;

            const dislikeBtn=document.createElement('img'); dislikeBtn.src="assets/icons/dislike.svg"; dislikeBtn.className='news-like-sm zoom-image';
            dislikeBtn.onclick=()=>likeNews(item.ID, userId, dislike);
            const spanDisliked=document.createElement('span'); spanDisliked.className='badge text-bg-secondary';spanDisliked.textContent=item.disliked;

            const viewBtn=document.createElement('img'); viewBtn.src="assets/icons/view.svg"; viewBtn.className='news-like-sm zoom-image';
            viewBtn.onclick=()=>viewNews(item.ID, userId);
            const spanViewed=document.createElement('span'); spanViewed.className='badge text-bg-secondary';spanViewed.textContent=item.viewCount;

            contentDiv.appendChild(title);
            contentDiv.appendChild(messageDate);
            contentDiv.appendChild(message);
            contentDiv.appendChild(type);
            contentDiv.appendChild(mediaLink);
            contentDiv.appendChild(downloadLink);
            contentDiv.appendChild(statusId);
            contentDiv.appendChild(newsId)

            contentDiv.appendChild(saveBtn);
            contentDiv.appendChild(deleteBtn);

            divLiked.appendChild(likeBtn);
            divLiked.appendChild(spanLiked);
            divLiked.appendChild(space);
            divLiked.appendChild(dislikeBtn);
            divLiked.appendChild(spanDisliked);
            divLiked.appendChild(space);
            divLiked.appendChild(viewBtn);
            divLiked.appendChild(spanViewed);
            div.appendChild(contentDiv);
            div.appendChild(divLiked);
            newsContainer.appendChild(div);
        });
    }catch(err){ newsContainer.innerHTML='<div>Ошибка загрузки новостей</div>'; }
}

function fillSelectStatus(selectElement){
    let option = document.createElement('option');
    option.value = "0";
    option.textContent = "Нет";
    selectElement.appendChild(option);

    option = document.createElement('option');
    option.value = "1";
    option.textContent = "Да";
    selectElement.appendChild(option);
}

async function saveNews(id,title,messageDate,message,typeStr,mediaLink,downloadLink,statusIdStr,authorId,liked,disliked,viewCount){
    const type = Number(typeStr);
    const statusId = Number(statusIdStr);

    try{
        const res=await fetch('https://localhost:8081/api/v1/news/'+id,{
            method:'PUT',
            headers:{'Content-Type':'application/json'},
            body:JSON.stringify({id,title,messageDate,message,type,mediaLink,downloadLink,statusId,authorId,liked,disliked,viewCount})
        });
        if (!res.ok) {
            const message = `An error occurred: ${res.status}`;
            throw new Error(message);
        }
        await loadNews();
    }catch(error)
    {
        console.error('Fetch error:', error);
        alert('Ошибка при сохранении');
        throw error; // Re-throw if necessary
    }
}

async function deleteNews(id){
    if(!confirm('Удалить новость?')) return;
    try{
        const res=await fetch('https://localhost:8081/api/v1/news/'+id,{
            method:'DELETE',
            headers:{'Content-Type':'application/json'},
            body:JSON.stringify({id})
        });
        if(!res.ok) throw new Error();
        await loadNews();
    }catch{ alert('Ошибка при удалении'); }
}

async function likeNews(newsId,userId,type){
    try{
        const res=await fetch('https://localhost:8081/api/v1/news/like/',{
            method:'PUT',
            headers:{'Content-Type':'application/json'},
            body:JSON.stringify({newsId,userId,type})
        });
        if (!res.ok) {
            const message = `An error occurred: ${res.status}`;
            throw new Error(message);
        }
        await loadNews();
    }catch(error)
    {
        console.error('Fetch error:', error);
        alert('Ошибка при сохранении');
        throw error; // Re-throw if necessary
    }
}

async function viewNews(newsId,userId){
    try{
        const res=await fetch('https://localhost:8081/api/v1/news/view/',{
            method:'PUT',
            headers:{'Content-Type':'application/json'},
            body:JSON.stringify({newsId,userId})
        });
        if (!res.ok) {
            const message = `An error occurred: ${res.status}`;
            throw new Error(message);
        }
        await loadNews();
    }catch(error)
    {
        console.error('Fetch error:', error);
        alert('Ошибка при сохранении');
        throw error; // Re-throw if necessary
    }
}

async function createNews(){
    const submitBtn = document.getElementById('submitBtn');

    // Отправка новости на сервер
    const title = document.getElementById('title').value.trim();
    const messageDate = document.getElementById('messageDate').value;
    const message = htmlEncode(document.getElementById('message').value.trim());
    const type = Number(document.getElementById('type').value);
    const mediaLink = document.getElementById('mediaLink').value.trim();
    const downloadLink = document.getElementById('downloadLink').value.trim();
    const statusId = Number(document.getElementById('statusId').value);
    const authorId = 0;

    if(!title) {
        alert('Заполните title');
        return;
    }
    if(!messageDate) {
        alert('Заполните messageDate');
        return;
    }
    if(!message) {
        alert('Заполните message');
        return;
    }
    if(!type) {
        alert('Заполните type');
        return;
    }

    const payload = { title, messageDate, message, type, mediaLink, downloadLink, statusId, authorId};

    submitBtn.disabled = true;
    try {
        const res = await fetch('https://localhost:8081/api/v1/news/', {
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify(payload)
        });
        if(!res.ok) throw new Error();
        alert('Новость успешно добавлена');
        // Очистка формы
        document.getElementById('title').value = '';
        document.getElementById('messageDate').value = '';
        document.getElementById('message').value = '';
        document.getElementById('type').value = '';
        document.getElementById('mediaLink').value = '';
        document.getElementById('downloadLink').value = '';
        document.getElementById('statusId').value = '';
    } catch(err) {
        alert('Ошибка при отправке новости на сервер');
    }
    submitBtn.disabled = false;
}

function htmlEncode(str) {
    const el = document.createElement('textarea');
    el.innerText = str;
    return el.innerHTML;
}

function decodeHtmlEntities(encodedString) {
    var textArea = document.createElement('textarea');
    textArea.innerHTML = encodedString;
    return textArea.value;
}