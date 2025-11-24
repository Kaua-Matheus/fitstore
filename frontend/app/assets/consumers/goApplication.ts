export async function Default() {
    const baseUrl = "http://localhost:8080";

    let response = await fetch(baseUrl + "/");
    if (!response.ok) {
        throw Error(`GET / failed`);
    } else {
        return response.json();
    }
}