function save() {
    username = document.getElementById('username')
    password = document.getElementById('password')

    let config = {
        username: username.value,
        password: password.value
    }
    chrome.storage.sync.set({ config: config });

}

let btn = document.getElementById('saveBtn')
btn.addEventListener('click', () => {
    save()
})