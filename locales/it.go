package locales

var It = map[string]interface{}{
	"address": map[string]interface{}{
		"default_country": []string{
			"Italia",
		},
		"city_prefix": []string{
			"San", "Borgo", "Sesto", "Quarto", "Settimo",
		},
		"city_suffix": []string{
			"a mare", "lido", "ligure", "del friuli", "salentino", "calabro", "veneto", "nell'emilia", "umbro", "laziale", "terme", "sardo",
		},
		"postcode": []string{
			"#####",
		},
		"state_abbr": []string{
			"AG", "AL", "AN", "AO", "AR", "AP", "AT", "AV", "BA", "BT", "BL", "BN", "BG", "BI", "BO", "BZ", "BS", "BR", "CA", "CL", "CB", "CI", "CE", "CT", "CZ", "CH", "CO", "CS", "CR", "KR", "CN", "EN", "FM", "FE", "FI", "FG", "FC", "FR", "GE", "GO", "GR", "IM", "IS", "SP", "AQ", "LT", "LE", "LC", "LI", "LO", "LU", "MC", "MN", "MS", "MT", "ME", "MI", "MO", "MB", "NA", "NU", "OT", "OR", "PD", "PA", "PR", "PV", "PG", "PU", "PE", "PC", "PI", "PT", "PN", "PZ", "PO", "RG", "RA", "RC", "RE", "RI", "RN", "RM", "RO", "SA", "VS", "SS", "SV", "SI", "SR", "SO", "TA", "TE", "TR", "TO", "OG", "TP", "TN", "TV", "TS", "UD", "VA", "VE", "VB", "VC", "VR", "VV", "VI", "VT",
		},
		"city": []string{
			"#{address.city_prefix} #{name.first_name} #{address.city_suffix}", "#{address.city_prefix} #{name.first_name}", "#{name.first_name} #{address.city_suffix}", "#{name.last_name} #{address.city_suffix}",
		},
		"street_name": []string{
			"#{address.street_suffix} #{name.first_name}", "#{address.street_suffix} #{name.last_name}",
		},
		"street_address": []string{
			"#{address.street_name} #{address.building_number}", "#{address.street_name} #{address.building_number}, #{secondary_address}",
		},
		"country": []string{
			"Afghanistan", "Albania", "Algeria", "American Samoa", "Andorra", "Angola", "Anguilla", "Antartide (territori a sud del 60° parallelo)", "Antigua e Barbuda", "Argentina", "Armenia", "Aruba", "Australia", "Austria", "Azerbaijan", "Bahamas", "Bahrain", "Bangladesh", "Barbados", "Bielorussia", "Belgio", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia", "Bosnia e Herzegovina", "Botswana", "Bouvet Island (Bouvetoya)", "Brasile", "Territorio dell'arcipelago indiano", "Isole Vergini Britanniche", "Brunei Darussalam", "Bulgaria", "Burkina Faso", "Burundi", "Cambogia", "Cameroon", "Canada", "Capo Verde", "Isole Cayman", "Repubblica Centrale Africana", "Chad", "Cile", "Cina", "Isola di Pasqua", "Isola di Cocos (Keeling)", "Colombia", "Comoros", "Congo", "Isole Cook", "Costa Rica", "Costa d'Avorio", "Croazia", "Cuba", "Cipro", "Repubblica Ceca", "Danimarca", "Gibuti", "Repubblica Dominicana", "Equador", "Egitto", "El Salvador", "Guinea Equatoriale", "Eritrea", "Estonia", "Etiopia", "Isole Faroe", "Isole Falkland (Malvinas)", "Fiji", "Finlandia", "Francia", "Guyana Francese", "Polinesia Francese", "Territori Francesi del sud", "Gabon", "Gambia", "Georgia", "Germania", "Ghana", "Gibilterra", "Grecia", "Groenlandia", "Grenada", "Guadalupa", "Guam", "Guatemala", "Guernsey", "Guinea", "Guinea-Bissau", "Guyana", "Haiti", "Heard Island and McDonald Islands", "Città del Vaticano", "Honduras", "Hong Kong", "Ungheria", "Islanda", "India", "Indonesia", "Iran", "Iraq", "Irlanda", "Isola di Man", "Israele", "Italia", "Giamaica", "Giappone", "Jersey", "Giordania", "Kazakhstan", "Kenya", "Kiribati", "Korea", "Kuwait", "Republicca Kirgiza", "Repubblica del Laos", "Latvia", "Libano", "Lesotho", "Liberia", "Libyan Arab Jamahiriya", "Liechtenstein", "Lituania", "Lussemburgo", "Macao", "Macedonia", "Madagascar", "Malawi", "Malesia", "Maldive", "Mali", "Malta", "Isole Marshall", "Martinica", "Mauritania", "Mauritius", "Mayotte", "Messico", "Micronesia", "Moldova", "Principato di Monaco", "Mongolia", "Montenegro", "Montserrat", "Marocco", "Mozambico", "Myanmar", "Namibia", "Nauru", "Nepal", "Antille Olandesi", "Olanda", "Nuova Caledonia", "Nuova Zelanda", "Nicaragua", "Niger", "Nigeria", "Niue", "Isole Norfolk", "Northern Mariana Islands", "Norvegia", "Oman", "Pakistan", "Palau", "Palestina", "Panama", "Papua Nuova Guinea", "Paraguay", "Peru", "Filippine", "Pitcairn Islands", "Polonia", "Portogallo", "Porto Rico", "Qatar", "Reunion", "Romania", "Russia", "Rwanda", "San Bartolomeo", "Sant'Elena", "Saint Kitts and Nevis", "Saint Lucia", "Saint Martin", "Saint Pierre and Miquelon", "Saint Vincent and the Grenadines", "Samoa", "San Marino", "Sao Tome and Principe", "Arabia Saudita", "Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Slovenia", "Isole Solomon", "Somalia", "Sud Africa", "Georgia del sud e South Sandwich Islands", "Spagna", "Sri Lanka", "Sudan", "Suriname", "Svalbard & Jan Mayen Islands", "Swaziland", "Svezia", "Svizzera", "Siria", "Taiwan", "Tajikistan", "Tanzania", "Tailandia", "Timor-Leste", "Togo", "Tokelau", "Tonga", "Trinidad e Tobago", "Tunisia", "Turchia", "Turkmenistan", "Isole di Turks and Caicos", "Tuvalu", "Uganda", "Ucraina", "Emirati Arabi Uniti", "Regno Unito", "Stati Uniti d'America", "United States Minor Outlying Islands", "Isole Vergini Statunitensi", "Uruguay", "Uzbekistan", "Vanuatu", "Venezuela", "Vietnam", "Wallis and Futuna", "Western Sahara", "Yemen", "Zambia", "Zimbabwe",
		},
		"building_number": []string{
			"###", "##", "#",
		},
		"street_suffix": []string{
			"Piazza", "Strada", "Via", "Borgo", "Contrada", "Rotonda", "Incrocio",
		},
		"secondary_address": []string{
			"Appartamento ##", "Piano #",
		},
		"state": []string{
			"Agrigento", "Alessandria", "Ancona", "Aosta", "Arezzo", "Ascoli Piceno", "Asti", "Avellino", "Bari", "Barletta-Andria-Trani", "Belluno", "Benevento", "Bergamo", "Biella", "Bologna", "Bolzano", "Brescia", "Brindisi", "Cagliari", "Caltanissetta", "Campobasso", "Carbonia-Iglesias", "Caserta", "Catania", "Catanzaro", "Chieti", "Como", "Cosenza", "Cremona", "Crotone", "Cuneo", "Enna", "Fermo", "Ferrara", "Firenze", "Foggia", "Forlì-Cesena", "Frosinone", "Genova", "Gorizia", "Grosseto", "Imperia", "Isernia", "La Spezia", "L'Aquila", "Latina", "Lecce", "Lecco", "Livorno", "Lodi", "Lucca", "Macerata", "Mantova", "Massa-Carrara", "Matera", "Messina", "Milano", "Modena", "Monza e della Brianza", "Napoli", "Novara", "Nuoro", "Olbia-Tempio", "Oristano", "Padova", "Palermo", "Parma", "Pavia", "Perugia", "Pesaro e Urbino", "Pescara", "Piacenza", "Pisa", "Pistoia", "Pordenone", "Potenza", "Prato", "Ragusa", "Ravenna", "Reggio Calabria", "Reggio Emilia", "Rieti", "Rimini", "Roma", "Rovigo", "Salerno", "Medio Campidano", "Sassari", "Savona", "Siena", "Siracusa", "Sondrio", "Taranto", "Teramo", "Terni", "Torino", "Ogliastra", "Trapani", "Trento", "Treviso", "Trieste", "Udine", "Varese", "Venezia", "Verbano-Cusio-Ossola", "Vercelli", "Verona", "Vibo Valentia", "Vicenza", "Viterbo",
		},
	},
	"company": map[string]interface{}{
		"suffix": []string{
			"SPA", "e figli", "Group", "s.r.l.",
		},
		"buzzwords": [][]string{
			[]string{
				"Abilità", "Access", "Adattatore", "Algoritmo", "Alleanza", "Analizzatore", "Applicazione", "Approccio", "Architettura", "Archivio", "Intelligenza artificiale", "Array", "Attitudine", "Benchmark", "Capacità", "Sfida", "Circuito", "Collaborazione", "Complessità", "Concetto", "Conglomerato", "Contingenza", "Core", "Database", "Data-warehouse", "Definizione", "Emulazione", "Codifica", "Criptazione", "Firmware", "Flessibilità", "Previsione", "Frame", "framework", "Funzione", "Funzionalità", "Interfaccia grafica", "Hardware", "Help-desk", "Gerarchia", "Hub", "Implementazione", "Infrastruttura", "Iniziativa", "Installazione", "Set di istruzioni", "Interfaccia", "Soluzione internet", "Intranet", "Conoscenza base", "Matrici", "Matrice", "Metodologia", "Middleware", "Migrazione", "Modello", "Moderazione", "Monitoraggio", "Moratoria", "Rete", "Architettura aperta", "Sistema aperto", "Orchestrazione", "Paradigma", "Parallelismo", "Policy", "Portale", "Struttura di prezzo", "Prodotto", "Produttività", "Progetto", "Proiezione", "Protocollo", "Servizio clienti", "Software", "Soluzione", "Standardizzazione", "Strategia", "Struttura", "Successo", "Sovrastruttura", "Supporto", "Sinergia", "Task-force", "Finestra temporale", "Strumenti", "Utilizzazione", "Sito web", "Forza lavoro",
			}, []string{
				"adattiva", "avanzata", "migliorata", "assimilata", "automatizzata", "bilanciata", "centralizzata", "compatibile", "configurabile", "cross-platform", "decentralizzata", "digitalizzata", "distribuita", "piccola", "ergonomica", "esclusiva", "espansa", "estesa", "configurabile", "fondamentale", "orizzontale", "implementata", "innovativa", "integrata", "intuitiva", "inversa", "gestita", "obbligatoria", "monitorata", "multi-canale", "multi-laterale", "open-source", "operativa", "ottimizzata", "organica", "persistente", "polarizzata", "proattiva", "programmabile", "progressiva", "reattiva", "riallineata", "ricontestualizzata", "ridotta", "robusta", "sicura", "condivisibile", "stand-alone", "switchabile", "sincronizzata", "sinergica", "totale", "universale", "user-friendly", "versatile", "virtuale", "visionaria",
			}, []string{
				"24 ore", "24/7", "terza generazione", "quarta generazione", "quinta generazione", "sesta generazione", "asimmetrica", "asincrona", "background", "bi-direzionale", "biforcata", "bottom-line", "coerente", "coesiva", "composita", "sensibile al contesto", "basta sul contesto", "basata sul contenuto", "dedicata", "didattica", "direzionale", "discreta", "dinamica", "eco-centrica", "esecutiva", "esplicita", "full-range", "globale", "euristica", "alto livello", "olistica", "omogenea", "ibrida", "impattante", "incrementale", "intangibile", "interattiva", "intermediaria", "locale", "logistica", "massimizzata", "metodica", "mission-critical", "mobile", "modulare", "motivazionale", "multimedia", "multi-tasking", "nazionale", "neutrale", "nextgeneration", "non-volatile", "object-oriented", "ottima", "ottimizzante", "radicale", "real-time", "reciproca", "regionale", "responsiva", "scalabile", "secondaria", "stabile", "statica", "sistematica", "sistemica", "tangibile", "terziaria", "uniforme", "valore aggiunto",
			},
		},
		"bs": [][]string{
			[]string{
				"partnerships", "comunità", "ROI", "soluzioni", "e-services", "nicchie", "tecnologie", "contenuti", "supply-chains", "convergenze", "relazioni", "architetture", "interfacce", "mercati", "e-commerce", "sistemi", "modelli", "schemi", "reti", "applicazioni", "metriche", "e-business", "funzionalità", "esperienze", "webservices", "metodologie",
			}, []string{
				"implementate", "utilizzo", "integrate", "ottimali", "evolutive", "abilitate", "reinventate", "aggregate", "migliorate", "incentivate", "monetizzate", "sinergizzate", "strategiche", "deploy", "marchi", "accrescitive", "target", "sintetizzate", "spedizioni", "massimizzate", "innovazione", "guida", "estensioni", "generate", "exploit", "transizionali", "matrici", "ricontestualizzate",
			}, []string{
				"valore aggiunto", "verticalizzate", "proattive", "forti", "rivoluzionari", "scalabili", "innovativi", "intuitivi", "strategici", "e-business", "mission-critical", "24/7", "globali", "B2B", "B2C", "granulari", "virtuali", "virali", "dinamiche", "magnetiche", "web", "interattive", "sexy", "back-end", "real-time", "efficienti", "front-end", "distributivi", "estensibili", "mondiali", "open-source", "cross-platform", "sinergiche", "out-of-the-box", "enterprise", "integrate", "di impatto", "wireless", "trasparenti", "next-generation", "cutting-edge", "visionari", "plug-and-play", "collaborative", "olistiche", "ricche",
			},
		},
		"name": []string{
			"#{name.last_name} #{name.suffix}", "#{name.last_name}-#{name.last_name} #{name.suffix}", "#{name.last_name}, #{name.last_name} e #{name.last_name} #{name.suffix}",
		},
	},
	"internet": map[string]interface{}{
		"free_email": []string{
			"gmail.com", "yahoo.com", "hotmail.com", "email.it", "libero.it", "yahoo.it",
		},
		"domain_suffix": []string{
			"com", "com", "com", "net", "org", "it", "it", "it",
		},
	},
	"name": map[string]interface{}{
		"first_name": []string{
			"Aaron", "Akira", "Alberto", "Alessandro", "Alighieri", "Amedeo", "Amos", "Anselmo", "Antonino", "Arcibaldo", "Armando", "Artes", "Audenico", "Ausonio", "Bacchisio", "Battista", "Bernardo", "Boris", "Caio", "Carlo", "Cecco", "Cirino", "Cleros", "Costantino", "Damiano", "Danny", "Davide", "Demian", "Dimitri", "Domingo", "Dylan", "Edilio", "Egidio", "Elio", "Emanuel", "Enrico", "Ercole", "Ermes", "Ethan", "Eusebio", "Evangelista", "Fabiano", "Ferdinando", "Fiorentino", "Flavio", "Fulvio", "Gabriele", "Gastone", "Germano", "Giacinto", "Gianantonio", "Gianleonardo", "Gianmarco", "Gianriccardo", "Gioacchino", "Giordano", "Giuliano", "Graziano", "Guido", "Harry", "Iacopo", "Ilario", "Ione", "Italo", "Jack", "Jari", "Joey", "Joseph", "Kai", "Kociss", "Laerte", "Lauro", "Leonardo", "Liborio", "Lorenzo", "Ludovico", "Maggiore", "Manuele", "Mariano", "Marvin", "Matteo", "Mauro", "Michael", "Mirco", "Modesto", "Muzio", "Nabil", "Nathan", "Nick", "Noah", "Odino", "Olo", "Oreste", "Osea", "Pablo", "Patrizio", "Piererminio", "Pierfrancesco", "Piersilvio", "Priamo", "Quarto", "Quirino", "Radames", "Raniero", "Renato", "Rocco", "Romeo", "Rosalino", "Rudy", "Sabatino", "Samuel", "Santo", "Sebastian", "Serse", "Silvano", "Sirio", "Tancredi", "Terzo", "Timoteo", "Tolomeo", "Trevis", "Ubaldo", "Ulrico", "Valdo", "Neri", "Vinicio", "Walter", "Xavier", "Yago", "Zaccaria", "Abramo", "Adriano", "Alan", "Albino", "Alessio", "Alighiero", "Amerigo", "Anastasio", "Antimo", "Antonio", "Arduino", "Aroldo", "Arturo", "Augusto", "Avide", "Baldassarre", "Bettino", "Bortolo", "Caligola", "Carmelo", "Celeste", "Ciro", "Costanzo", "Dante", "Danthon", "Davis", "Demis", "Dindo", "Domiziano", "Edipo", "Egisto", "Eliziario", "Emidio", "Enzo", "Eriberto", "Erminio", "Ettore", "Eustachio", "Fabio", "Fernando", "Fiorenzo", "Folco", "Furio", "Gaetano", "Gavino", "Gerlando", "Giacobbe", "Giancarlo", "Gianmaria", "Giobbe", "Giorgio", "Giulio", "Gregorio", "Hector", "Ian", "Ippolito", "Ivano", "Jacopo", "Jarno", "Joannes", "Joshua", "Karim", "Kris", "Lamberto", "Lazzaro", "Leone", "Lino", "Loris", "Luigi", "Manfredi", "Marco", "Marino", "Marzio", "Mattia", "Max", "Michele", "Mirko", "Moreno", "Nadir", "Nazzareno", "Nestore", "Nico", "Noel", "Odone", "Omar", "Orfeo", "Osvaldo", "Pacifico", "Pericle", "Pietro", "Primo", "Quasimodo", "Radio", "Raoul", "Renzo", "Rodolfo", "Romolo", "Rosolino", "Rufo", "Sabino", "Sandro", "Sasha", "Secondo", "Sesto", "Silverio", "Siro", "Tazio", "Teseo", "Timothy", "Tommaso", "Tristano", "Umberto", "Ariel", "Artemide", "Assia", "Azue", "Benedetta", "Bibiana", "Brigitta", "Carmela", "Cassiopea", "Cesidia", "Cira", "Clea", "Cleopatra", "Clodovea", "Concetta", "Cosetta", "Cristyn", "Damiana", "Danuta", "Deborah", "Demi", "Diamante", "Diana", "Donatella", "Doriana", "Edvige", "Elda", "Elga", "Elsa", "Emilia", "Enrica", "Erminia", "Eufemia", "Evita", "Fatima", "Felicia", "Filomena", "Flaviana", "Fortunata", "Gelsomina", "Genziana", "Giacinta", "Gilda", "Giovanna", "Giulietta", "Grazia", "Guendalina", "Helga", "Ileana", "Ingrid", "Irene", "Isabel", "Isira", "Ivonne", "Jelena", "Jole", "Claudia", "Kayla", "Kristel", "Laura", "Lucia", "Lia", "Lidia", "Lisa", "Loredana", "Loretta", "Luce", "Lucrezia", "Luna", "Maika", "Marcella", "Maria", "Mariagiulia", "Marianita", "Mariapia", "Marieva", "Marina", "Maristella", "Maruska", "Matilde", "Mecren", "Mercedes", "Mietta", "Miriana", "Miriam", "Monia", "Morgana", "Naomi", "Nayade", "Nicoletta", "Ninfa", "Noemi", "Nunzia", "Olimpia", "Oretta", "Ortensia", "Penelope", "Piccarda", "Prisca", "Rebecca", "Rita", "Rosalba", "Rosaria", "Rosita", "Ruth", "Samira", "Sarita", "Selvaggia", "Shaira", "Sibilla", "Soriana", "Thea", "Tosca", "Ursula", "Vania", "Vera", "Vienna", "Violante", "Vitalba", "Zelida",
		},
		"last_name": []string{
			"Amato", "Barbieri", "Barone", "Basile", "Battaglia", "Bellini", "Benedetti", "Bernardi", "Bianc", "Bianchi", "Bruno", "Caputo", "Carbon", "Caruso", "Cattaneo", "Colombo", "Cont", "Conte", "Coppola", "Costa", "Costantin", "D'amico", "D'angelo", "Damico", "De Angelis", "De luca", "De rosa", "De Santis", "Donati", "Esposito", "Fabbri", "Farin", "Ferrara", "Ferrari", "Ferraro", "Ferretti", "Ferri", "Fior", "Fontana", "Galli", "Gallo", "Gatti", "Gentile", "Giordano", "Giuliani", "Grassi", "Grasso", "Greco", "Guerra", "Leone", "Lombardi", "Lombardo", "Longo", "Mancini", "Marchetti", "Marian", "Marini", "Marino", "Martinelli", "Martini", "Martino", "Mazza", "Messina", "Milani", "Montanari", "Monti", "Morelli", "Moretti", "Negri", "Neri", "Orlando", "Pagano", "Palmieri", "Palumbo", "Parisi", "Pellegrini", "Pellegrino", "Piras", "Ricci", "Rinaldi", "Riva", "Rizzi", "Rizzo", "Romano", "Ross", "Rossetti", "Ruggiero", "Russo", "Sala", "Sanna", "Santoro", "Sartori", "Serr", "Silvestri", "Sorrentino", "Testa", "Valentini", "Villa", "Vitale", "Vitali",
		},
		"prefix": []string{
			"Sig.", "Dott.", "Dr.", "Ing.",
		},
		"suffix": []string{},
		"name": []string{
			"#{name.prefix} #{name.first_name} #{name.last_name}", "#{name.first_name} #{name.last_name}", "#{name.first_name} #{name.last_name}", "#{name.first_name} #{name.last_name}", "#{name.first_name} #{name.last_name}", "#{name.first_name} #{name.last_name}",
		},
	},
	"phone_number": map[string]interface{}{
		"formats": []string{
			"+## ### ## ## ####", "+## ## #######", "+## ## ########", "+## ### #######", "+## ### ########", "+## #### #######", "+## #### ########", "0## ### ####", "+39 0## ### ###", "3## ### ###", "+39 3## ### ###"}}}
