//
//  ResultView.swift
//  GitGraph BFS
//
//  Created by Ryan e Letícia on 18/09/23.
//

import SwiftUI

struct ResultView: View {
    var paths: [String]
    var baconNumber: Int
    
    var body: some View {
        VStack(alignment: .leading, spacing: 16) {
            Text("A menor ligação entre os usuários é:")
            Text("Numero de Bacon: \(baconNumber)")
            
            Text("Caminho:")
            
            List(paths, id: \.self) {
                path in
                    Text(path)
                        .font(.headline)
            }
        }.navigationTitle("Resultado").padding()
        .preferredColorScheme(ColorScheme.dark)
        
    }
}

struct ResultView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationStack {
            ResultView(paths: ["Leticia", "Ryan"], baconNumber: 2)
        }
    }
}
	
