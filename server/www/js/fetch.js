export function fetch_get(url) {
    const cData = fetch(url)
        .then((resp) => resp.json())
        .then((data) => {
            console.log("fetch_get:", data);
            return data;
        });

    //   const data = async () => {
    //     const data = await cData;
    //     console.log(data);
    //     return data;
    //   };

    return cData;
}