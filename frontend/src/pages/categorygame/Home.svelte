<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Button_custom from '../../components/button_custom.svelte'
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel_info from '../../components/Panel_info.svelte' 
    import Panel from '../../components/Panel_default.svelte' 

    export let path_api = "";
    export let font_size = "";
    export let token = "";
    export let listHome = [];
    export let totalrecord = 0;

    let page = "Category Game";
    let sData = "New";
    let isModal_Form_New = false
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_flag = false;
    let buttonLoading_class = "";
    let msg_error = "";
    let searchHome = "";
    let filterHome = [];
    let categame_id_field = ""
    let categame_id_enabled = true;
    let categame_display_field = 0;
    let creategame_create_field = ""
    let creategame_update_field = ""
    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        categame_id_field: yup
            .string()
            .required("ID is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "ID must Character A-Z or a-z or 1-9"
            )
            .min(2, "ID must be at least 4 Character")
            .max(5, "ID must be at most 5 Character"),
        categame_name_field: yup
            .string()
            .required("Name is Required")
            .matches(
                /^[a-zA-z0-9\- ]+$/,
                "Name must Character A-Z or a-z or 1-9"
            )
            .min(4, "Name must be at least 4 Character")
            .max(70, "Name must be at most 70 Character"),
         categame_status_field: yup.string().required("Status is Required"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            categame_id_field: "",
            categame_name_field: "",
            categame_status_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(
                values.categame_id_field,
                values.categame_name_field,
                values.categame_status_field
            );
        },
    });
    async function SaveTransaksi(idcate, name, status) {
        let flag = true;
        msg_error = "";
        if(status == ""){
            flag = false;
            msg_error += "The Status is required";
        }
        if (flag) {
            buttonLoading_flag = true;
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecategorygame", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "CATEGAME-SAVE",
                    idcategame: idcate.toUpperCase(),
                    name: name,
                    display: parseInt(categame_display_field),
                    status: status,
                }),
            });
            const json = await res.json();
            if(!res.ok){
                loader_msg = "System Mengalami Trouble"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
            }else{
                if (json.status == 200) {
                    loader_msg = json.message
                    if(sData == "New"){
                        $form.categame_id_field= "";
                        $form.categame_name_field= "";
                        $form.categame_status_field= "";
                        categame_display_field = 0;
                    }
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_flag = false;
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                RefreshHalaman();
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true;
            }
        }
    }
    
    const EntryData = (tipeentry,id,name,display,status,create,update) => {
        if(tipeentry == "Edit"){
            sData = "Edit"
            categame_id_field = id
            $form.categame_id_field = id;
            $form.categame_name_field = name;
            $form.categame_status_field = status;
            categame_display_field = display;
            creategame_create_field = create;
            creategame_update_field = update;
            categame_id_enabled = false;
            isModal_Form_New = true;
        }else{
            sData = "New"
            clearField()
            categame_id_enabled = true;
            isModal_Form_New = true;
        }
        
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        EntryData("New","","","","")
        
    };
    
    function clearField(){
        $form.categame_id_field = "";
        $form.categame_name_field = "";
        $form.categame_status_field = "";
        $errors.categame_id_field = "";
        $errors.categame_name_field = "";
        $errors.categame_status_field = "";
        categame_id_field = ""
        categame_display_field = 0;
        creategame_update_field = ""
        creategame_create_field = ""
    }
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_nama
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) 
            );
        } else {
            filterHome = [...listHome];
        }
    }
</script>
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />

<Panel
    on:handleNewForm={NewData}
    on:handleRefresh={RefreshHalaman}
    panel_button_new={true}
    panel_button_refresh={true}
    panel_page="{page}"
    panel_total="{totalrecord}">
    <slot:template slot="panel_search">
        <div class="absolute inset-y-0 left-0 flex items-center pl-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
        </div>
        <input 
            bind:value={searchHome}
            type="text" placeholder="Search by Nama" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center"></th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">NO</th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">STATUS</th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-left">ID</th>
                    <th width="*" class="bg-[#475289] {font_size} text-white text-left">NAMA</th>
                    <th width="7%" class="bg-[#475289] {font_size} text-white text-right">DISPLAY</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td on:click={() => {
                            EntryData("Edit",rec.home_id,rec.home_nama,rec.home_display,rec.home_status,rec.home_create,rec.home_update);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </td>
                        <td class="{font_size} align-top text-center">{rec.home_no}</td>
                        <td class="{font_size} align-top text-center"><span class="{rec.home_statusclass} text-center rounded-md p-1 px-2">{rec.home_status}</span></td>
                        <td class="{font_size} align-top text-left">{rec.home_id}</td>
                        <td class="{font_size} align-top text-left">{rec.home_nama}</td>
                        <td class="{font_size} align-top text-right text-blue-500 font-semibold">{rec.home_display}</td>
                    </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="10" class="text-center">
                            <progress class="self-start progress progress-primary w-56"></progress>
                        </td>
                    </tr>
                </tbody>
            {/if}
        </table>
    </slot:template>
</Panel>


<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_New}>
<Modal_popup
    modal_popup_id="my-modal-formnew"
    modal_popup_title="Entry/{sData}"
    modal_popup_class="select-none max-w-full lg:max-w-xl overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-auto flex-col overflow-auto gap-3 mt-2 ">
            <div class="mt-2">
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_maxlength_text="5"
                    input_invalid={$errors.categame_id_field.length > 0}
                    bind:value={$form.categame_id_field}
                    input_id="categame_id_field"
                    input_enabled={categame_id_enabled}
                    input_text_class="uppercase"
                    input_placeholder="ID"/>
                {#if $errors.categame_id_field}
                    <small class="text-pink-600 text-[11px]">{$errors.categame_id_field}</small>
                {/if}
            </div>
            <div >
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_maxlength_text="70"
                    input_invalid={$errors.categame_name_field.length > 0}
                    bind:value={$form.categame_name_field}
                    input_id="categame_name_field"
                    input_placeholder="Nama"/>
                {#if $errors.categame_name_field}
                    <small class="text-pink-600 text-[11px]">{$errors.categame_name_field}</small>
                {/if}
            </div>
            <Input_custom
                input_enabled={true}
                input_tipe="number_nolabel"
                bind:value={categame_display_field}
                input_maxlenght="3"
                input_id="categame_display_field"
                input_placeholder="Display"/>
            <div class="">
                <select
                    on:change="{handleChange}"
                    bind:value={$form.categame_status_field}
                    invalid={$errors.categame_status_field.length > 0} 
                    class="w-full text-sm lg:text-md rounded px-3  border border-gray-300 focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none">
                    <option disabled selected value="">--Pilih Status--</option>
                    <option value="ACTIVE">ACTIVE</option>
                    <option value="DEACTIVE">DEACTIVE</option>
                </select>
                {#if $errors.categame_status_field}
                    <small class="text-pink-600 text-[11px]">{$errors.categame_status_field}</small>
                {/if}
            </div>
            {#if sData == "Edit"}
                <Panel_info>
                    <slot:template slot="panel_body">
                        <table>
                            <tr>
                                <td>Create</td>
                                <td>:</td>
                                <td>{creategame_create_field}</td>
                            </tr>
                            {#if creategame_update_field != ""}
                            <tr>
                                <td>Modified</td>
                                <td>:</td>
                                <td>{creategame_update_field}</td>
                            </tr>
                            {/if}
                        </table>
                    </slot:template>
                </Panel_info>
            {/if}
        </div>
        <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
            <Button_custom 
                on:click={() => {
                    handleSubmit();
                }}
                button_disable={buttonLoading_flag}
                button_class=""
                button_disable_class="{buttonLoading_class}"
                button_title="Submit" />
        </div>
    </slot:template>
</Modal_popup>



<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



