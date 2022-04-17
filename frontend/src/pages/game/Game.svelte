<script>
    import Home from "../game/Home.svelte";
    import Modal_alert from '../../components/Modal_alert.svelte' 
    export let path_api = ""
    export let font_size = "";
    let listHome = [];
    let listcategorygame = [];
    let listprovidergame = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let akses_page = false;
    let isModalNotif = false;
    let msg_error = ""
    async function initapp() {
        msg_error = ""
        const res = await fetch(path_api+"api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "GAME-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            msg_error = json.message;
            akses_page = false;
        } else {
            akses_page = true;
            initHome();
        }
        if(msg_error != ""){
            isModalNotif = true;
        }
    }
    async function initHome() {
        listHome = []
        listcategorygame = []
        listprovidergame = []
        const res = await fetch(path_api+"api/game", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
            }),
        });
        const json = await res.json();
        if (!res.ok) {
			isModalNotif = true;
            msg_error = "Maaf Saat Ini Anda Tidak Bisa Mengakses Halaman Ini"
		}else{
            if (json.status == 200) {
                record = json.record;
                let status_class = "";
                let status_text = "";
                let no = 0;
                let recordlistcategorygame = json.listcategorygame;
                let recordlistprovidergame = json.listprovidergame;
                if (record != null) {
                    totalrecord = record.length;
                    for (var i = 0; i < record.length; i++) {
                        no = no + 1;
                        if (record[i]["providergame_status"] == "Y") {
                            status_class = "bg-[#ebfbee] text-[#6ec07b]"
                            status_text = "ACTIVE"
                        } else {
                            status_class = "bg-[#fde3e3] text-[#ea7779]"
                            status_text = "DEACTIVE"
                        }
                        listHome = [
                            ...listHome,
                            {
                                home_no: no,
                                home_id: record[i]["providergame_id"],
                                home_nama: record[i]["providergame_name"],
                                home_phone: record[i]["providergame_phone"],
                                home_email: record[i]["providergame_email"],
                                home_note: record[i]["providergame_note"],
                                home_status: status_text,
                                home_create: record[i]["providergame_create"],
                                home_update: record[i]["providergame_update"],
                                home_statusclass: status_class,
                            },
                        ];
                    }
                }
                if (recordlistcategorygame != null) {
                    for (let i = 0; i < recordlistcategorygame.length; i++) {
                        
                        listcategorygame = [
                            ...listcategorygame,
                            {
                                categame_id:recordlistcategorygame[i]["categame_id"],
                                categame_name:recordlistcategorygame[i]["categame_name"],
                            },
                        ];
                    }
                }
                if (recordlistprovidergame != null) {
                    for (let i = 0; i < recordlistprovidergame.length; i++) {
                        listprovidergame = [
                            ...listprovidergame,
                            {
                                providergame_id:recordlistprovidergame[i]["providergame_id"],
                                providergame_name:recordlistprovidergame[i]["providergame_name"],
                            },
                        ];
                    }
                }
            }else {
                isModalNotif = true;
                msg_error = "Maaf Sistem Sedang Mengalami Masalah"
            }
        }
        setTimeout(function () {
            isModalNotif = false;
        }, 1000);
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 1000);
    };
    const handleLogout = (e) => {
        logout()
    }
    initapp();
</script>
{#if akses_page == true}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleLogout={handleLogout}
        {path_api}
        {listcategorygame}
        {listprovidergame}
        {font_size}
        {token}
        {listHome}
        {totalrecord}/>
{/if}
<input type="checkbox" id="my-modal-notiffirst" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notiffirst" 
    modal_title="Information" modal_message="{msg_error}"  />