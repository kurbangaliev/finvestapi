async function authenticate() {
    const existingToken = localStorage.getItem('jwt');
    const login = document.getElementById('login').value;
    const password = document.getElementById('password').value;
    const message = document.getElementById('message');
    const btnLogin = document.getElementById('btnLogin');

    if (!login || !password) {
        message.textContent = 'Введите логин и пароль';
        message.style.color = 'red';
        return;
    }

    try {
        btnLogin.disabled = true;
        const response = await fetch('https://localhost:8081/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({ login, password })
        });

        if (!response.ok) {
            throw new Error('Ошибка авторизации');
        }

        const result = await response.json();

        // ожидаем, что сервер вернёт JWT
        const token = result.token;
        if (token) {
            // сохраняем JWT
            localStorage.setItem('jwt', token);
        }

        message.textContent = 'Успешный вход';
        message.style.color = 'green';
        console.log('Ответ сервера:', result);

        setTimeout(() => {
            window.location.href = '/newsBrowser';
        }, 800);

    } catch (err) {
        message.textContent = 'Неверный логин или пароль';
        message.style.color = 'red';
        btnLogin.disabled = false;
    }
}