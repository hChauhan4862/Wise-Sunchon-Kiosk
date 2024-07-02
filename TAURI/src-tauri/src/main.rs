// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use tauri::Manager;

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tauri::command]
async fn open_app_window(app: tauri::AppHandle) {
    tauri::window::WindowBuilder::new(
        &app,
        "main_window",
        tauri::WindowUrl::App("index.html".into()),
    )
    .title("Wise Neosco Kiosk Application")
    .fullscreen(true)
    .resizable(false)
    .build()
    .unwrap();
}

#[tauri::command]
async fn serial_string(message: String, app: tauri::AppHandle) {
    app.emit_all("SERIAL_EVENT", Payload { message: message.into() }).unwrap();
}

#[tauri::command]
fn get_server_url() -> String {
    // let url = "http://13.233.128.10:8080/V1/KIOSK".into();
    let url = "http://localhost:8080/V1/KIOSK".into();
    return url;
}

// the payload type must implement `Serialize` and `Clone`.
#[derive(Clone, serde::Serialize)]
struct Payload {
    message: String,
}

fn main() {
    
    tauri::Builder::default()
        .setup(|app| {
            let handle = app.handle();
            std::thread::spawn(move || {
                _ = tauri::WindowBuilder::new(
                    &handle,
                    "home",
                    tauri::WindowUrl::App("home.html".into()),
                )
                .title("Wise Kiosk")
                .center()
                .inner_size(225.0, 215.0)
                .resizable(false)
                .build();
            });
            Ok(())
        })
        .invoke_handler(tauri::generate_handler![open_app_window, get_server_url, serial_string])
        .run(tauri::generate_context!())
        .expect("error while running tauri application")
}
