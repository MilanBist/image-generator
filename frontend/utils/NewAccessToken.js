export async function getNewAccessToken(refreshToken){
    console.log("Refresh token is: ", refreshToken);
    let tokenBody = {
        refreshToken: refreshToken,
    }

    try{
        const resp = await apiClient.post("/refreshToken", tokenBody);
        let responseData = resp.data["message"];
        if(responseData === "expired"){
            return 401;
        }
        localStorage.setItem("accessToken", resp.data["data"]["accessToken"]);
        return 200;
    }catch(err){
        console.log("Error in generatin the new refresh token: ", err);
        return 400;
    }
}