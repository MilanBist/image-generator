import { useState, useEffect } from "react";
import apiClient from "../utils/Base";
import { useNavigate } from "react-router-dom";
import { getNewAccessToken } from "../utils/NewAccessToken";


// async function getHistory(setHistory){
//     const navigate = useNavigate();
//     let accessToken = localStorage.getItem("accessToken");
//     let refreshToken = localStorage.getItem("refreshToken");
//     try{
//         const resp = await apiClient.get("getHistory", {
//             headers:{
//                 Authorization: `Bearer ${accessToken}`,
//             }
//         })

//         // if successfull then just return
//         if (resp.data["success"] === true){
//             // just return the data from here
//             setHistory(resp.data["data"]);
//             return 200;
//         }
//     }catch(err){
//         const errorStatus = err.response.status;
//         if (errorStatus ===  401){
//             // get new access token
//             const response = await getNewAccessToken(refreshToken);
//             if (response === 401){
//                 return 401;
//             }

//             if (response === 200){
//                 // successfully obtained new access token now you can again call the function of get history
//                 getHistory();
//                 return 200;
//             }
//         }
//         return 400;
//     }
    
// }

function History({output, history, setHistory}){
    const navigate = useNavigate();
    // async function getHistory() {
    // }

    useEffect(() => {
        let refreshToken = localStorage.getItem("refreshToken");
        let accessToken = localStorage.getItem("accessToken");

        if(accessToken === null || refreshToken === null){
            alert("You are not logged in. \n Redirecting to login page.");
            return;
        }
        // let response = getHistory(setHistory);
        // if (response !== 200){
        //     alert("You are not authorized. Please login again.");
        //     navigate("/");
        //     return;
        // }
    }, [output]);

    return (
        <>
            <div className="history-page">
            </div>

        </>
    )
}

export default History;