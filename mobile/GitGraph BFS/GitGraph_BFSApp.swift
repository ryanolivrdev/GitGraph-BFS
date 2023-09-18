//
//  GitGraph_BFSApp.swift
//  GitGraph BFS
//
//  Created by Ryan e Letícia on 17/09/23.
//

import SwiftUI

@main
struct GitGraph_BFSApp: App {
    @State private var path: NavigationPath = NavigationPath()

    var body: some Scene {
        WindowGroup {
            NavigationStack(path: $path) {
                ContentView()
            }
        }
    }
}
