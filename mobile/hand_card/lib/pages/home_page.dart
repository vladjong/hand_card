import 'dart:convert';
import 'package:hand_card/model/card.dart';
import 'package:http/http.dart' as http;


import 'package:flutter/material.dart';
class HomePage extends StatelessWidget {
  HomePage(this.jwt) {}

  final String jwt;

  // Future<List<DiscountCard>> cards = getCards();

  Future<List<DiscountCard>> getCards() async {
    print('Bearer $jwt');
    final url = 'http://10.0.2.2:8080/api/cards';
    final queryCoord = {
      "lat": 82.897918,
      "lon": 54.980332
    };
    
    // final uri = Uri.parse(url).replace(queryParameters: queryCoord);
    final response = await http.get(
      Uri.parse(url),
      headers: {
        'Content-Type': 'application/json; charset=UTF-8',
        'Authorization': 'Bearer $jwt',
      },
    );
    print(response.body);
    final body = json.decode(response.body);
    return body.map<DiscountCard>(DiscountCard.fromJson).toList();
  }

  @override
  Widget build(BuildContext context) =>
    Scaffold(
      body: SafeArea(
          child: Column(
            children: [
              SizedBox(height: 10,),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 25.0),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  // ignore: prefer_const_literals_to_create_immutables
                  children: [
                    // ignore: prefer_const_literals_to_create_immutables
                    Row(children: [
                      const Text(
                      'Мои',
                      style: TextStyle(
                        fontSize: 28,
                        color: Colors.deepPurpleAccent,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    const Text(
                      ' карты',
                      style: TextStyle(fontSize: 28),
                    ),
                    ],),
                    ElevatedButton(
                      onPressed: (){},
                      style: ElevatedButton.styleFrom(
                        shape: CircleBorder(),
                        padding: EdgeInsets.all(7),
                        backgroundColor: Colors.deepPurpleAccent,
                        foregroundColor: Colors.purple,
                      ),
                      child: const Icon(Icons.add, color: Colors.white,),
                    ),
                  ]
                ),
              ),
              SizedBox(height: 25,),
              FutureBuilder<List<DiscountCard>>(
                future: getCards(),
                builder: (context, snapshot) {
                  if (snapshot.hasData) {
                    final cards = snapshot.data!;
                    return buildCards(cards);
                  } else {
                    return const Text("No cards data");
                  }
                }
              )
            ]
          ),
      )
    );
          // FutureBuilder<List<DiscountCard>>(
          // future: getCards(),
          // builder: (context, snapshot) {
          //   if (snapshot.hasData) {
          //     final cards = snapshot.data!;
          //     return buildCards(cards);
          //   } else {
          //     return const Text("No cards data");
          //   }
          // },
          // )

  Widget buildCards(List<DiscountCard> cards) => ListView.builder(
    itemCount: cards.length,
    shrinkWrap: true,
    itemBuilder: (context, index) {
      final card = cards[index];
      return Card(
        color: Colors.grey,
        child: Padding(
          padding: const EdgeInsets.all(14),
          child: ListTile(
            title: Text(card.organization),
            trailing: Container(
              width: 15,
              child: Row(
                children: [
                  Expanded(
                    child: IconButton(
                      onPressed:() {
                       cards.removeAt(index);
                    }, icon: Icon(Icons.delete, color: Colors.white,)))
                ],)
            ),
          ),
        ),
              // title: Text(card.organization)    // subtitle: Text(card.number),
      );
    }
  );

}