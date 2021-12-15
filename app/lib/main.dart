import 'package:flutter/material.dart';
import 'package:controlcenter/lenovo.dart';

void main() {
  runApp(const Application());
}

class Application extends StatelessWidget {
  const Application({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Control Center',
      theme: ThemeData(
        primarySwatch: Colors.deepPurple,
      ),
      initialRoute: '/lenovo',
      routes: <String, WidgetBuilder>{
        '/lenovo': (BuildContext context) =>
            const LenovoSettingsPage(title: 'Lenovo'),
      },
    );
  }
}
