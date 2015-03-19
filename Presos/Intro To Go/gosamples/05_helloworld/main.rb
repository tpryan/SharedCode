greet_list = { 
			'Chinese' => "你好世界",
			'Dutch' => "Hello wereld",
			'English' => "Hello world",
			'French' => "Bonjour monde",
			'German' => "Hallo Welt",
			'Greek' => "γειά σου κόσμος",
			'Italian' => "Ciao mondo",
			'Japanese' => "こんにちは世界",
			'Korean' => "여보세요 세계",
			'Portuguese' =>  "Olá mundo",
			'Russian' => "Здравствулте мир",
			'Spanish' => "Hola mundo"
}	
lang = ARGV[0]
lang = "English" if lang == nil
response = greet_list[lang]

if response == nil
	puts "The language you selected is not valid."
else
	puts greet_list[lang]
end

