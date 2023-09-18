//
//  ContentView.swift
//  GitGraph BFS
//
//  Created by Ryan e Letícia on 17/09/23.
//

import SwiftUI

struct ResultModel: Codable {
    var path: [String]
    var kevinBaconNumber: Int
}

struct ContentView: View {
    private let urlApi = "http://localhost:8080"
    @State private var userOrigin: String = ""
    @State private var userTarget: String = ""
    
    @State private var paths: [String] = [""]
    @State private var baconNumber: Int = 0
    @State private var loading: Bool = false
    @State private var showAlert = false
    @State private var viewResult = false

    func fetchData() async {
        let url = URL(string: urlApi + "?userOrigin=\(userOrigin)&userTarget=\(userTarget)")!
        
        do {
            let (data, response) = try await URLSession.shared.data(from: url)
            
            guard let httpResponse = response as? HTTPURLResponse,
                  (200...299).contains(httpResponse.statusCode) else {
                print("Erro na resposta.")
                showAlert = true
                
                return
            }
            
            if let decodedResponse = try? JSONDecoder().decode(ResultModel.self, from: data) {
                paths = decodedResponse.path
                baconNumber = decodedResponse.kevinBaconNumber
                viewResult = true
            }
        } catch {
            print("These data are not valid")
        }
    }
    
    var body: some View {
        ZStack {
            Color("BackgroundColor")
                .ignoresSafeArea()
            
            VStack(spacing: 16) {
                Text("GitGraph BFS")
                    .foregroundColor(.white)
                    .font(.title)
                    .bold()
                
                Text("Busque a menor ligação entre dois usuários do GitHub")
                    .foregroundColor(.white)
                    .multilineTextAlignment(.center)
                
                VStack(alignment: .leading) {
                    Text("Usuário de origem")
                        .foregroundColor(.white)
                    
                    TextField("Joe Doe", text: $userOrigin)
                        .textInputAutocapitalization(.never)
                        .foregroundColor(.black)
                        .frame(maxWidth: .infinity)
                        .padding(10)
                        .background(.white)
                        .cornerRadius(5.0)
                        .disabled(loading)
                    
                    Spacer().frame(height: 16)
                    
                    Text("Usuário de destino")
                        .foregroundColor(.white)
                    
                    TextField("Joe Doe", text: $userTarget)
                        .textInputAutocapitalization(.never)
                        .foregroundColor(.black)
                        .frame(maxWidth: .infinity)
                        .padding(10)
                        .background(.white)
                        .cornerRadius(5.0)
                        .disabled(loading)
                }
                
                Button {
                    Task {
                        loading = true
                        await fetchData()
                        loading = false
                    }
                } label: {
                    Text(!loading ? "Realizar Buscar" : "Carregando...")
                        .bold()
                        .foregroundColor(.white)
                        .padding(10)
                        .frame(maxWidth: .infinity)
                        .background(!loading ? Color("ButtonColor") : .gray)
                        .cornerRadius(5.0)
                    
                }
                .alert("Error", isPresented: $showAlert) {
                    Button("OK") { }
                } message: {
                    Text("Nao foi possivel realizar a busca, verifique se os usuarios estao corretos e tente novamente.")
                }
                
                
            }.navigationDestination(
                isPresented: $viewResult) {
                    ResultView(paths: paths, baconNumber: baconNumber)
                }.padding().preferredColorScheme(ColorScheme.dark)
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationStack {
            ContentView()
        }
    }
}
