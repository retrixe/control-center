import 'package:dbus/dbus.dart';
import 'package:flutter/material.dart';

abstract class SettingsPage extends StatelessWidget {
  const SettingsPage({Key? key, required this.title, required this.client})
      : super(key: key);

  final String title;
  final DBusClient client;

  Widget buildPage(BuildContext context);

  @override
  Widget build(BuildContext context) {
    const horizontalDivider = Divider(
      height: 1,
      indent: 0,
      endIndent: 0,
    );

    return Scaffold(
      appBar: AppBar(
        title: const Text('Linux Control Center'),
      ),
      body: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(
            width: 160,
            child: ListView(
              scrollDirection: Axis.vertical,
              shrinkWrap: true,
              children: [
                ListTile(
                  selected: title == 'Lenovo',
                  leading: const Icon(Icons.laptop),
                  title: const Text('Lenovo'),
                  onTap: () {
                    Navigator.of(context).pushReplacementNamed('/lenovo');
                  },
                ),
                horizontalDivider,
                ListTile(
                  selected: title == 'nouveau',
                  leading: const Icon(Icons.highlight_alt),
                  title: const Text('nouveau'),
                  onTap: () {
                    Navigator.of(context).pushReplacementNamed('/nouveau');
                  },
                ),
                horizontalDivider,
              ],
            ),
          ),
          const VerticalDivider(
            width: 1,
            indent: 0,
            endIndent: 0,
          ),
          Expanded(child: buildPage(context)),
        ],
      ),
    );
  }

  Future<void> showDBusError(BuildContext context, List<Widget> text) async {
    // The navigator is never poppable unless there's a dialog.
    if (Navigator.of(context).canPop()) return;
    return showDialog<void>(
      context: context,
      barrierDismissible: false,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Error'),
          content: SingleChildScrollView(
            child: ListBody(
              children: text,
            ),
          ),
          actions: <Widget>[
            TextButton(
              child: const Text('Dismiss'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }
}
