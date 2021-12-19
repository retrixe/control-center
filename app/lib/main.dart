import 'package:control_center/lenovo.dart';
import 'package:control_center/nouveau.dart';
import 'package:dbus/dbus.dart';
import 'package:flutter/material.dart';
import 'package:yaru/yaru.dart';

void main() => runApp(const Application());

class Application extends StatefulWidget {
  const Application({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() => _ApplicationState();
}

class _ApplicationState extends State<Application> {
  late DBusClient client;

  @override
  void initState() {
    super.initState();
    client = DBusClient.system();
  }

  @override
  void dispose() {
    super.dispose();
    client.close().ignore();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Linux Control Center',
      theme: yaruLight, // ThemeData(primarySwatch: Colors.deepPurple),
      darkTheme: yaruDark,
      initialRoute: '/lenovo',
      routes: <String, WidgetBuilder>{
        '/lenovo': (BuildContext context) => LenovoSettingsPage(client: client),
        '/nouveau': (BuildContext context) =>
            NouveauSettingsPage(client: client),
      },
    );
  }
}
