export async function authenticate(_currentState: any, formData: FormData) {
    console.log(formData.get('username'))
    console.log(formData.get('password'))
    const response = await fetch('http://localhost:8082/auth', {
        method: 'POST',
        body: JSON.stringify({
            "username": formData.get('username'),
            "password": formData.get('password')
            })
    });

    if (response.status === 200) {
        window.location.href = '/';
        console.log('success')
    }
}
