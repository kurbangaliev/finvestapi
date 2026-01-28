async function logout() {
    await fetch('/logout', {
        method: 'POST',
        credentials: 'include'
    });
    window.location.href = '/';
}

async function authFetch(url, options = {}) {
    alert('authFetch');
    const response = await fetch(url, {
        ...options,
        credentials: "include"
    });

    if (response.status === 401) {
        window.location.href = "/login";
        throw new Error("Unauthorized");
    }

    return response;
}